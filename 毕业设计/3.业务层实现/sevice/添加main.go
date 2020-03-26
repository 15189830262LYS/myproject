[......]

    // 根据证书编号与名称查询信息
    result, err := serviceSetup.FindEduByCertNoAndName("222","李四")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        var edu service.Education
        json.Unmarshal(result, &edu)
        fmt.Println("根据证书编号与姓名查询信息成功：")
        fmt.Println(edu)
    }

    //===========================================//

}




[......]

    // 根据身份证号码查询信息
    result, err = serviceSetup.FindEduInfoByEntityID("101")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        var edu service.Education
        json.Unmarshal(result, &edu)
        fmt.Println("根据身份证号码查询信息成功：")
        fmt.Println(edu)
    }

    //===========================================//

}




[......]

    // 修改/添加信息
    info := service.Education{
        Name: "张三",
        Gender: "男",
        Nation: "汉",
        EntityID: "101",
        Place: "北京",
        BirthDay: "1991年01月01日",
        EnrollDate: "2013年9月",
        GraduationDate: "2015年7月",
        SchoolName: "中国政法大学",
        Major: "社会学",
        QuaType: "普通",
        Length: "两年",
        Mode: "普通全日制",
        Level: "研究生",
        Graduation: "毕业",
        CertNo: "333",
        Photo: "/static/phone/11.png",
    }
    msg, err = serviceSetup.ModifyEdu(info)
    if err != nil {
        fmt.Println(err.Error())
    }else {
        fmt.Println("信息操作成功, 交易编号为: " + msg)
    }

    //===========================================//

}



[......]

    // 根据身份证号码查询信息
    result, err = serviceSetup.FindEduInfoByEntityID("101")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        var edu service.Education
        json.Unmarshal(result, &edu)
        fmt.Println("根据身份证号码查询信息成功：")
        fmt.Println(edu)
    }

    //===========================================//

}



[......]

    // 根据证书编号与名称查询信息
    result, err = serviceSetup.FindEduByCertNoAndName("333","张三")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        var edu service.Education
        json.Unmarshal(result, &edu)
        fmt.Println("根据证书编号与姓名查询信息成功：")
        fmt.Println(edu)
    }

    //===========================================//

}