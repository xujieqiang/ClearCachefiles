package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type Filesconf struct {
	Temp   string `json:"temp_files"`
	Wechat string `json:"wechat_files"`
}

func main() {
	fmt.Println("清理缓存和临时文件！\n………………\n\n")
	// a := path.Join("./", "a a", "aa.txt")
	// aba, _ := filepath.Abs(a)
	wechat, temp, err := Getfiles()
	if err != nil {
		fmt.Println(err)
		return
	}
	roming, _ := os.LookupEnv("APPDATA")
	adobe_path := path.Join(roming, "Adobe", "Common")
	_, err = os.ReadDir(adobe_path)
	if err != nil {

		fmt.Println("Adobe/common 文件夹不存在，不用清理！")
	} else {
		fmt.Println("\n\n========================")
		fmt.Println("  查找Adobe的缓存文件夹！")
		fmt.Println("========================")
		fmt.Println(adobe_path)
		fmt.Println()
		err := Delfile(adobe_path)
		if err != nil {
			fmt.Println(err)
		}
	}
	//fmt.Println("====>", wechat)

	wechat_path := path.Join(wechat, "Documents", "WeChat Files")

	abwechat, _ := filepath.Abs(wechat_path)
	_, err = os.ReadDir(abwechat)
	if err != nil {

		fmt.Println("WeChat Files 文件夹不存在，不用清理！")
	} else {
		fmt.Println("\n\n========================")
		fmt.Println("  查找微信的缓存文件夹！")
		fmt.Println("========================")
		fmt.Println(wechat_path)
		fmt.Println()
		err := DelEverything(wechat_path)
		if err != nil {
			fmt.Println(err)
		}
	}

	temp_path := temp
	fmt.Println("\n\n========================")
	fmt.Println("  查找temp临时文件夹！")
	fmt.Println("========================")
	fmt.Println(temp_path)
	fmt.Println()
	err2 := DelEverything(temp_path)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("已经清理完毕！")
	fmt.Scanln()

	// f, err := os.Open(aba)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(errors.New("不能读取带空格的文件名！"))

	// }
	// defer f.Close()
	// b, err := io.ReadAll(f)
	// if err != nil {
	// 	fmt.Println(errors.New("读取文件内容时发生错误！"))

	// }
	// fmt.Println(string(b))

}

func Getfiles() (string, string, error) {
	// var cfg Filesconf
	// files, err := os.Open("config.json")
	// if err != nil {
	// 	//fmt.Println(err, "无法读取文件！config.json")
	// 	return "", "", err
	// }
	// defer files.Close()
	// content, err := io.ReadAll(files)
	// if err != nil {
	// 	//fmt.Println(err)
	// 	return "", "", err
	// }
	// fmt.Println(string(content))
	// //data := os.ExpandEnv(string(content))
	// //fmt.Println(data)
	// err = json.Unmarshal(content, &cfg)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", "", err
	// }
	//fmt.Println(cfg.Temp, cfg.Wechat)
	/**********************************************
	读取文件完成
	**************************************/
	wechat, _ := os.LookupEnv("USERPROFILE")
	temp, _ := os.LookupEnv("TEMP")
	return wechat, temp, nil
}

func DelEverything(files string) error {

	//newpath := path.Join(files, "WeChat Files")
	//fmt.Println("newpath==============>", newpath)

	// ee := os.Chdir(files)
	// if ee != nil {
	// 	return errors.New("下面的路径无法切换：" + files)
	// }
	fs, err := os.ReadDir(files)
	if err != nil {
		fmt.Println(errors.New("不能读取这样的路径：" + files))
		return err
	}

	for _, val := range fs {
		if val.IsDir() {

			dirpath := path.Join(files, val.Name())
			//abpath, _ := filepath.Abs(dirpath)
			DelEverything(dirpath)
			os.Remove(dirpath)

			fmt.Println(dirpath, "被删除！")
		} else {

			dirpath := path.Join(files, val.Name())
			//abpath, _ := filepath.Abs(dirpath)
			rr := os.Remove(dirpath)
			if rr != nil {
				return rr
			}
			fmt.Println(dirpath, "被删除！")
		}
	}
	return nil
}

func Delfile(files string) error {
	fs, err := os.ReadDir(files)
	if err != nil {
		fmt.Println(errors.New("不能读取这样的路径：" + files))
		return err
	}
	for _, val := range fs {
		if val.IsDir() {
			dirpath := path.Join(files, val.Name())
			Delfile(dirpath)
		} else {
			filepath := path.Join(files, val.Name())
			rr := os.Remove(filepath)
			if rr != nil {
				return rr
			}
			fmt.Println(filepath, " 被删除！")
		}
	}
	return nil
}

// func DelTemp(files string) error {
// 	// arr := strings.Split(files, string(os.PathSeparator))
// 	// cd := len(arr)
// 	// last := arr[cd-1]

// 	// ee := os.Chdir(files)
// 	// if ee != nil {
// 	// 	return ee
// 	// }
// 	fs, err := os.ReadDir(files)
// 	if err != nil {
// 		return err
// 	}

// 	for _, val := range fs {
// 		if val.IsDir() {

// 			dirpath := path.Join(files, val.Name())
// 			abpath, _ := filepath.Abs(dirpath)
// 			DelTemp(abpath)

// 			os.Remove(abpath)
// 			fmt.Println(abpath, "被删除！")
// 		} else {

// 			dirpath := path.Join(files, val.Name())
// 			abpath, _ := filepath.Abs(dirpath)
// 			os.Remove(abpath)
// 			fmt.Println(abpath, "被删除！")
// 		}
// 	}

// 	return nil
// }

// func Delfiles_wechat(files string) error {
// 	_, err := os.ReadDir(files)
// 	if err != nil {
// 		return err
// 	}

// 	newpath := path.Join(files, "Documents", "WeChat Files")
// 	_, err = os.ReadDir(newpath)
// 	if err != nil {
// 		return err
// 	}
// 	abpath, _ := filepath.Abs(newpath)

// 	err = DelWechat(abpath)
// 	if err != nil {
// 		return err
// 	}

// 	// err = os.Remove(abpath)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// fmt.Println(abpath, "被删除！")

// 	return nil
// }

// func Delfiles_temp(files string) error {
// 	fs, err := os.ReadDir(files)
// 	if err != nil {
// 		return err
// 	}
// 	for _, val := range fs {
// 		if val.IsDir() {

// 			dirpath := path.Join(files, val.Name())

// 			abpath, _ := filepath.Abs(dirpath)

// 			DelTemp(abpath)

// 			os.Remove(abpath)
// 			fmt.Println(abpath, "被删除！")

// 		} else {
// 			dirpath := path.Join(files, val.Name())
// 			abpath, _ := filepath.Abs(dirpath)
// 			os.Remove(abpath)
// 			fmt.Println(abpath, "被删除！")

// 		}
// 	}
// 	return nil
// }
