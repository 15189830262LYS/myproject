import(
    [......]
    "github.com/kongyixueyuan.com/education/web/controller"
    "github.com/kongyixueyuan.com/education/web"
)

func main(){}
    [......]

    app := controller.Application{
        Setup: &serviceSetup,
    }
    web.WebStart(app)
}