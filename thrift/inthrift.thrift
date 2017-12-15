namespace php In

service ApiService {

    string   ping();
    string   call(1:string service_name,2:string method,3:string params='',4:string request_info="")
}