package main

import (
    
    "github.com/go-martini/martini"
    "github.com/yosssi/ace"
    "github.com/yosssi/martini-acerender"
)

func main() {

    m := martini.Classic()
    //加载acerender
    m.Use(acerender.Renderer(&acerender.Options{
      AceOptions: &ace.Options{
        //定义模板路径
        BaseDir: "templates",
      },
    }))   

    //定义静态文件路径
    m.Use(martini.Static("bower_components"))
    m.Use(martini.Static("components"))

    m.Get("/", func(r acerender.Render) {

      r.HTML(200, "index", nil, nil)
    })

    m.Run()
}
