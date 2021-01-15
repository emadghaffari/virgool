# production example

### key:
confs

### value:

{"Environment":"production","GRPC":{"Host":"localhost","Port":":8082","Endpoint":":8083"},"HTTP":{"Host":"localhost","Port":":8080","Endpoint":":8081"},"DEBUG":{"Host":"localhost","Port":":8084","Endpoint":":8085"},"Redis":{"Username":"","Password":"","DB":0,"Host":"redis:6379","Logger":false},"Vault":{"Address":"http://vault:8200","Token":"s.9JFm7dyhXVIagWhPEUSiYTAN"},"MultiInstanceMode":false,"Log":{"disable_colors":false,"quote_empty_fields":false},"Service":{"Name":"notification","MinCL":1000,"MaxCl":9999,"Redis":{"SMSDuration":300000000000,"SMSCodeVerification":10000000000,"UserDuration":40000000000000}},"Jaeger":{"HostPort":"jaeger:6831","LogSpans":true},"Kafka":{"Username":"","Password":"","Brokers":["kafka1:9092","kafka2:9092","kafka3:9092"],"Version":"v1","Group":"","Assignor":"range","Oldest":true,"Verbose":false,"Topics":{"Notif":"notifications"},"Auth":false,"Consumer":true,"Producer":false},"Notif":{"SMS":{"UserAPIKey":"dd198b76e5eaa968sda98sd41d1ef31f8b76","SecretKey":"cp6teBC!@FeB6a5sd1C!YFuBC","Token":{"URL":"https://RestfulSms.com/api/Token","ContentType":"application/json"},"Send":{"TemplateURL":"https://RestfulSms.com/api/UltraFastSend","Verify":{"TemplateID":"22108","ContentType":"application/json"}}}}}