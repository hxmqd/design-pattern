package main

import "fmt"

type Image interface {
	display()
}

type RealImage struct {
	fileName string
}

func (r *RealImage)display(){
	fmt.Println("Displaying "+r.fileName)
}

func (r *RealImage)loadFromDisk(fileName string){
	fmt.Println("Loading "+ fileName)
}

func NewRealImage(fileName string) *RealImage{
	realImage := &RealImage{
		fileName: fileName,
	}
	realImage.loadFromDisk(fileName)
	return realImage
}

type ProxyImage struct {
	realImage *RealImage
	fileName string
}

func (r *ProxyImage)display(){
	if r.realImage == nil {
		r.realImage = NewRealImage(r.fileName)
	}
	r.realImage.display()
}

func NewProxyImage(fileName string) *ProxyImage{
	return &ProxyImage{fileName: fileName}
}

func main(){
	var image Image
	image = NewProxyImage("test.jpg")
	// 首次图像从磁盘加载
	image.display()
	fmt.Println("")
	// 图像无需从磁盘加载
	image.display()

}


