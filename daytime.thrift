namespace go api

struct DaytimeRequest {
  1: string clientDaytime
}

struct DaytimeResponse {
  1: string serverDaytime
}

service service_name_daytime_in_thrift_file{
    DaytimeResponse daytime(1: DaytimeRequest req)
}
