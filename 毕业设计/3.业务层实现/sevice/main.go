package main

import (
    [......]
    "github.com/kongyixueyuan.com/education/service"
)

[......]
    //===========================================//

    serviceSetup := service.ServiceSetup{
        ChaincodeID:EduCC,
        Client:channelClient,
    }

    edu := service.Education{
        Name: "张三",
        Gender: "男",
        Nation: "汉",
        EntityID: "101",
        Place: "北京",
        BirthDay: "1991年01月01日",
        EnrollDate: "2009年9月",
        GraduationDate: "2013年7月",
        SchoolName: "中国政法大学",
        Major: "社会学",
        QuaType: "普通",
        Length: "四年",
        Mode: "普通全日制",
        Level: "本科",
        Graduation: "毕业",
        CertNo: "111",
        Photo: "/static/phone/11.png",
    }

    edu2 := service.Education{
        Name: "李四",
        Gender: "男",
        Nation: "汉",
        EntityID: "102",
        Place: "上海",
        BirthDay: "1992年02月01日",
        EnrollDate: "2010年9月",
        GraduationDate: "2014年7月",
        SchoolName: "中国人民大学",
        Major: "行政管理",
        QuaType: "普通",
        Length: "四年",
        Mode: "普通全日制",
        Level: "本科",
        Graduation: "毕业",
        CertNo: "222",
        Photo: "/static/phone/22.png",
    }

    msg, err := serviceSetup.SaveEdu(edu)
    if err != nil {
        fmt.Println(err.Error())
    }else {
        fmt.Println("信息发布成功, 交易编号为: " + msg)
    }

    msg, err = serviceSetup.SaveEdu(edu2)
    if err != nil {
        fmt.Println(err.Error())
    }else {
        fmt.Println("信息发布成功, 交易编号为: " + msg)
    }    

    //===========================================//

}