# production example

### key:
confs

### value:
{"Environment":"production","GRPC":{"Host":"localhost","Port":":8082","Endpoint":":8083"},"HTTP":{"Host":"localhost","Port":":8080","Endpoint":":8081"},"DEBUG":{"Host":"localhost","Port":":8084","Endpoint":":8085"},"MYSQL":{"Username":"root","Password":"password","Host":"db","Schema":"virgool","Driver":"mysql","Automigrate":true,"Logger":true,"Namespace":""},"Redis":{"Username":"","Password":"","DB":0,"Host":"redis:6379","Logger":false},"Vault":{"Address":"http://vault:8200","Token":"s.9JFm7dyhXVIagWhPEUSiYTAN"},"MultiInstanceMode":false,"Log":{"disable_colors":false,"quote_empty_fields":false},"Service":{"Name":"auth","Redis":{"SMSDuration":1000000000,"SMSCodeVerification":1000000000,"UserDuration":1000000000}},"Jaeger":{"HostPort":"jaeger:6831","LogSpans":true},"Kafka":{"Username":"admin","Password":"admin-secret","Brokers":["kafka1:9092","kafka2:9092","kafka3:9092"],"Version":"v1","Group":"","Assignor":"range","Oldest":true,"Verbose":false,"Topics":{"Notif":"notifications"},"Auth":false,"Consumer":false,"Producer":true},"JWT":{"RSecret":"yyyyyyyyyyyyyyy","Secret":"xxxxxxxxxxxxx"}}