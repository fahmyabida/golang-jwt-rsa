package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

func main(){
	//key := generateRSA()
	

//	sToken := ` eyJhbGciOiJSUzI1NiIsImtpZCI6IjEifQ.eyJoZWxsbyI6I
//mZhaG15In0.XCSWf38wWjr3fftQbh52wEtjk1L5rmblKgg_PtZQd9I1JynCJ5h2q
//KqJEYtV0l87cakmaVbWeIfFzxEcadwAYhEcbfh8McfQs-P_JsFqNC4FaUhORkNb1
//Ytfv1aBniy_hRC4rnoDnpFkXsPBNbCrdK0mSOPuaF0RJAKrloXYy9-K5WIXZ7yGv
//t2jZwj3WCOh85bHLM8XmjKEBu4zCsO0z6pG7LWKACyixVRMAdg9fiMoZmqVphhhO
//V9LC3nYvjZ1cvruRLCVtELMcRLQGCH23HKFPml_MnxpzJjYOXwPTGl4bdwGJGgmb
//7jdD0Wyc6Lv9QhZ7TJwjjfGLtG-sLmRmA`
//
//	rsaPub := `-----BEGIN PUBLIC KEY-----
//MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv
//vkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc
//aT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy
//tvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0
//e+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb
//V6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9
//MwIDAQAB
//-----END PUBLIC KEY-----`
//
//	rsaPriv := `-----BEGIN RSA PRIVATE KEY-----
//MIIEogIBAAKCAQEAnzyis1ZjfNB0bBgKFMSvvkTtwlvBsaJq7S5wA+kzeVOVpVWw
//kWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHcaT92whREFpLv9cj5lTeJSibyr/Mr
//m/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIytvHWTxZYEcXLgAXFuUuaS3uF9gEi
//NQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0e+lf4s4OxQawWD79J9/5d3Ry0vbV
//3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWbV6L11BWkpzGXSW4Hv43qa+GSYOD2
//QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9MwIDAQABAoIBACiARq2wkltjtcjs
//kFvZ7w1JAORHbEufEO1Eu27zOIlqbgyAcAl7q+/1bip4Z/x1IVES84/yTaM8p0go
//amMhvgry/mS8vNi1BN2SAZEnb/7xSxbflb70bX9RHLJqKnp5GZe2jexw+wyXlwaM
//+bclUCrh9e1ltH7IvUrRrQnFJfh+is1fRon9Co9Li0GwoN0x0byrrngU8Ak3Y6D9
//D8GjQA4Elm94ST3izJv8iCOLSDBmzsPsXfcCUZfmTfZ5DbUDMbMxRnSo3nQeoKGC
//0Lj9FkWcfmLcpGlSXTO+Ww1L7EGq+PT3NtRae1FZPwjddQ1/4V905kyQFLamAA5Y
//lSpE2wkCgYEAy1OPLQcZt4NQnQzPz2SBJqQN2P5u3vXl+zNVKP8w4eBv0vWuJJF+
//hkGNnSxXQrTkvDOIUddSKOzHHgSg4nY6K02ecyT0PPm/UZvtRpWrnBjcEVtHEJNp
//bU9pLD5iZ0J9sbzPU/LxPmuAP2Bs8JmTn6aFRspFrP7W0s1Nmk2jsm0CgYEAyH0X
//+jpoqxj4efZfkUrg5GbSEhf+dZglf0tTOA5bVg8IYwtmNk/pniLG/zI7c+GlTc9B
//BwfMr59EzBq/eFMI7+LgXaVUsM/sS4Ry+yeK6SJx/otIMWtDfqxsLD8CPMCRvecC
//2Pip4uSgrl0MOebl9XKp57GoaUWRWRHqwV4Y6h8CgYAZhI4mh4qZtnhKjY4TKDjx
//QYufXSdLAi9v3FxmvchDwOgn4L+PRVdMwDNms2bsL0m5uPn104EzM6w1vzz1zwKz
//5pTpPI0OjgWN13Tq8+PKvm/4Ga2MjgOgPWQkslulO/oMcXbPwWC3hcRdr9tcQtn9
//Imf9n2spL/6EDFId+Hp/7QKBgAqlWdiXsWckdE1Fn91/NGHsc8syKvjjk1onDcw0
//NvVi5vcba9oGdElJX3e9mxqUKMrw7msJJv1MX8LWyMQC5L6YNYHDfbPF1q5L4i8j
//8mRex97UVokJQRRA452V2vCO6S5ETgpnad36de3MUxHgCOX3qL382Qx9/THVmbma
//3YfRAoGAUxL/Eu5yvMK8SAt/dJK6FedngcM3JEFNplmtLYVLWhkIlNRGDwkg3I5K
//y18Ae9n7dHVueyslrb6weq7dTkYDi3iOYRW8HRkIQh06wEdbxt0shTzAJvvCQfrB
//jg/3747WSsf/zBTcHihTRBdAv6OmdhV4/dD5YBfLAkLrd+mX7iE=
//-----END RSA PRIVATE KEY-----`


	bPrvKey,_ := ioutil.ReadFile("private2.pem")
	pKey := BytesToPrivateKey(bPrvKey)
	bPubKey,_ := ioutil.ReadFile("public2.pem")
	pubKey := BytesToPublicKey(bPubKey)

	//t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	//t.Claims = &CustomClaimsExample{
	//	&jwt.StandardClaims{
	//		// set the expire time
	//		// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
	//		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
	//	},
	//	"level1",
	//	CustomerInfo{user, "human"},
	//}

	// Creat token string
	//return t.SignedString(signKey)

	mClaim := jwt.MapClaims{}
	mClaim["hello"] = "fahmy"


	algMethod := jwt.GetSigningMethod("RS256")
	t := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": algMethod.Alg(),
			"kid": 1, //should generated id
		},
		Claims: mClaim,
		Method: algMethod,
	}
	s, err:= t.SignedString(pKey)
	s2, err:= t.SignedString(pKey)
	s3, err:= t.SignedString(pKey)

	//string, err:= algMethod.Sign("this is it", pKey)
	fmt.Println(s, err)
	fmt.Println()
	fmt.Println(s2,err)
	fmt.Println()
	fmt.Println(s3,err)
	s = "eyJhbGciOiJSUzI1NiIsImtpZCI6MSwidHlwIjoiSldUIn0.eyJoZWxsbyI6ImZhaG15In0.WpE5t_0zTLdbrJrNNJQiZPJGd6Rmb24BEsNW539aYSMu_W4igknspdppVGYP0LYM8IApwib_3itWWEHlH9u6hrOEOgYLNVuAZCUcnvX7TF0Nzlf0e9XFNJUuefMW_6iujWiD4pQLskFvt7j3UPJN-YpxkAhxu6mS3acRxniu8-9hl2_6tVVy-nDFDthjlYaYWzkeMGx_lxV7Afzk5yGQTu0sUKpm3E4Gc-mI28YwjG81LzT3hVfex5v1G_nAX4IC9VqCKGCrmqV9Ct6Bu04Ts2PESbShbEocgdGztMUjBEmjsRE0IP-qq1eoCZT79qe3ZqjsBot_FL_KTVNGEPiZGw"

	token, err := jwt.Parse(s,func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return pubKey, nil
	})
	fmt.Println()
	fmt.Println(token.Valid)
	fmt.Println()
	fmt.Println(token, err)
	//mySigningKey := []byte("yolo")
	//sToken, errToken := token.SignedString(mySigningKey)


	//tokenByte, _ := claimKu.RSASign("RS256", pKey)
	//fmt.Println("generated token ", string(tokenByte))
	//
	//claims, err := jwt.RSACheck(tokenByte, pubKey)
	//fmt.Println("result ", claims, err)
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		fmt.Println(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		fmt.Println("not ok")
	}
	return key
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		fmt.Println(err)
	}
	return key
}