[......]

func (t *ServiceSetup) FindEduByCertNoAndName(certNo, name string) ([]byte, error){

    req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryEduByCertNoAndName", Args: [][]byte{[]byte(certNo), []byte(name)}}
    respone, err := t.Client.Query(req)
    if err != nil {
        return []byte{0x00}, err
    }

    return respone.Payload, nil
}



[......]

func (t *ServiceSetup) FindEduInfoByEntityID(entityID string) ([]byte, error){

    req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryEduInfoByEntityID", Args: [][]byte{[]byte(entityID)}}
    respone, err := t.Client.Query(req)
    if err != nil {
        return []byte{0x00}, err
    }

    return respone.Payload, nil
}




[......]

func (t *ServiceSetup) ModifyEdu(edu Education) (string, error) {

    eventID := "eventModifyEdu"
    reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
    defer t.Client.UnregisterChaincodeEvent(reg)

    // 将edu对象序列化成为字节数组
    b, err := json.Marshal(edu)
    if err != nil {
        return "", fmt.Errorf("指定的edu对象序列化时发生错误")
    }

    req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "updateEdu", Args: [][]byte{b, []byte(eventID)}}
    respone, err := t.Client.Execute(req)
    if err != nil {
        return "", err
    }

    err = eventResult(notifier, eventID)
    if err != nil {
        return "", err
    }

    return string(respone.TransactionID), nil
}