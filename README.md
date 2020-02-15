# golang并发爬虫
爬取zhenai网

### 项目结构：
src:
    data -- 数据存储相关  
    engine -- 爬虫引擎  
    fetcher -- 数据获取，网络请求  
    model  -- 数据实体  
    scheduler -- worker调度器  
    zhenai
        parser -- 页面解析
    main.go -- 程序入口
   
  ### 基础单任务版爬虫架构
  engine -----request--->fetcher  
  fetcher -----result---->parser  
  parser ----parseResult ---> engine  
  
  其中parseResult中会有新的request,这样就会形成一个图形的爬取结构  
  城市列表 ---->  [城市1,2,3] -----> 用户[1,....n]
  
  ### 并发版的爬虫架构
  并发版爬虫将 fetcher和parser还有engine的一部分抽象成一个worker, 
  并发得调度这个worker，通过channel 来传递数据，就做成一个并发的爬虫
  
    