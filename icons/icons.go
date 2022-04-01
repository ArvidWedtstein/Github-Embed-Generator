package icons

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type IconsSVG []struct {
	Name string `json:"name"`
	Svg  string `json:"svg"`
}

func Icons(icon string) string {
	jsonFile, err := os.Open("icons/icons.json")

	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var iconssvg IconsSVG
	json.Unmarshal(byteValue, &iconssvg)
	iconf := ""
	for _, v := range iconssvg {
		if v.Name == icon {
			iconf = v.Svg
		}
	}
	if iconf != "" {
		return iconf
	} else {
		return "<svg version='1.1' xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' width='20' height='20' x='10' y='10' viewBox='0 0 1920 1920' style='enable-background:new 0 0 1920 1920;' xml:space='preserve'><style type='text/css'>.st0{fill:#050505;}.st1{fill:#F5F12B;}.st2{fill:#FF0000;}</style><rect x='873.4' y='1533.1' transform='matrix(0.9295 0.3688 -0.3688 0.9295 669.8132 -243.1051)' class='st0' width='194.6' height='194.6'/><rect x='865.1' y='1417.7' transform='matrix(0.9295 0.3688 -0.3688 0.9295 626.6744 -248.1424)' class='st0' width='194.6' height='194.6'/><g><polygon class='st0' points='675.1,1260.3 494.2,1188.5 425.3,1361.9 422.4,1361.9 422.4,1556.6 617,1556.6 617,1543.1 683.5,1375.7 636.7,1357.1 	'/><polygon class='st0' points='1491,1253.9 1306.7,1191.5 1244.8,1374 1241.3,1374 1241.3,1568.7 1435.9,1568.7 1435.9,1555.2 1506.1,1378.3 1455.7,1358.3 	'/><polygon class='st0' points='1431.7,727.1 1498.2,559.7 1451.4,541.1 1489.8,444.3 1308.9,372.5 1240.1,545.9 1237.1,545.9 1237.1,740.6 1431.7,740.6 	'/><polygon class='st0' points='1015.8,527.5 1019.6,529 1090.5,350.4 900.7,275.1 871.2,349.6 458,349.6 458,531.2 1015.8,531.2 	'/><polygon class='st0' points='1575.2,918.6 1560.3,956 380.8,956 380.8,350.8 343.6,350.8 395.8,219.3 285.6,175.6 216.1,350.8 215.9,350.8 215.9,956 215.9,988 215.9,1150.7 821.1,1150.7 821.1,1756.8 1015.8,1756.8 1015.8,1150.7 1632.9,1150.7 1632.9,1149.2 1704.1,969.8 	'/></g><rect x='488.8' y='1182.7' class='st1' width='194.6' height='194.6'/><path class='st1' d='M1328.2,366.8h153.8c11.3,0,20.4,9.2,20.4,20.4V541c0,11.3-9.2,20.4-20.4,20.4h-153.8 c-11.3,0-20.4-9.2-20.4-20.4V387.3C1307.7,376,1316.9,366.8,1328.2,366.8z'/><path class='st1' d='M1309.5,1182.7h191.1c1,0,1.8,0.8,1.8,1.8v191.1c0,1-0.8,1.8-1.8,1.8h-191.1c-1,0-1.8-0.8-1.8-1.8v-191.1 C1307.7,1183.5,1308.5,1182.7,1309.5,1182.7z'/><polygon class='st2' points='484.1,776.3 484.1,355.5 1089.4,355.5 1089.4,163.2 277.6,163.2 277.6,350.8 277.6,355.5 277.6,967.9 484.1,967.9 899.5,967.9 899.5,1585.1 1094.1,1585.1 1094.1,967.9 1704.1,967.9 1704.1,776.3 '/></svg>"
	}
}
