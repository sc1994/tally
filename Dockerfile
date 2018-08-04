# 选择nginx服务
FROM nginx:latest
# copy代码
COPY . /src
# 添加nginx配置文件
COPY nginx.conf /etc/nginx/nginx.conf
# 去掉默认的nginx配置文件
RUN rm /etc/nginx/conf.d/default.conf