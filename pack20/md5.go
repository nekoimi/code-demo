/**
 * # ------------Oooo---
 * # -----------(----)---
 * # ------------)--/----
 * # ------------(_/-
 * # ----oooO----
 * # ----(---)----
 * # -----\--(--
 * # ------\_)-
 * # ----
 * #     author : Yprisoner <yyprisoner@gmail.com>
 * #     19-5-26 下午10:18
 * #                            ------
 * #    「 涙の雨が頬をたたくたびに美しく 」
 **/
package pack20

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func Pack20MD5()  {
	md5 := md5.New()
	md5.Write([]byte("demo"))
	result := md5.Sum([]byte(""))
	fmt.Printf("%x \n",result)
}

func Pack20SHA256()  {
	sha := sha256.New()
	sha.Write([]byte("demo"))
	result := sha.Sum([]byte(""))
	fmt.Printf("%x \n",result)
}
