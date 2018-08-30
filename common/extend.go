package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// BindExtend 扩展熟悉绑定方法防止400异常
func BindExtend(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindWith(obj, binding.JSON)
	return err
}

// clip 图片裁剪
// func clip(in io.Reader, out io.Writer, x0, y0, x1, y1, quality int) error {
// 	origin, fm, err := image.Decode(in)
// 	if err != nil {
// 		return err
// 	}

// 	switch fm {
// 	case "jpeg":
// 		img := origin.(*image.YCbCr)
// 		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
// 		return jpeg.Encode(out, subImg, &jpeg.Options{quality})
// 	case "png":
// 		switch canvas.(type) {
// 		case *image.NRGBA:
// 			img := canvas.(*image.NRGBA)
// 			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
// 			return png.Encode(out, subImg)
// 		case *image.RGBA:
// 			img := canvas.(*image.RGBA)
// 			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
// 			return png.Encode(out, subImg)
// 		}
// 	case "gif":
// 		img := origin.(*image.Paletted)
// 		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.Paletted)
// 		return gif.Encode(out, subImg, &gif.Options{})
// 	case "bmp":
// 		img := origin.(*image.RGBA)
// 		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
// 		return bmp.Encode(out, subImg)
// 	default:
// 		return errors.New("ERROR FORMAT")
// 	}
// 	return nil
// }
