FROM docker.io/centos

MAINTAINER AUTHOR'S LUODI FROM RODDYPY.COM QQ:923401910

WORKDIR /usr/local/src/

COPY pcre-8.37.tar.gz    .
COPY nginx-1.8.0.tar.gz  .

#add user www
RUN useradd -s /sbin/nologin www  && yum install gcc gcc-c++ openssl openssl-devel -y

RUN tar zxf pcre-8.37.tar.gz   &&  cd pcre-8.37 && ./configure --prefix=/usr/local/pcre  && make && make install

RUN tar zxf nginx-1.8.0.tar.gz && cd nginx-1.8.0 && ./configure --prefix=/usr/local/nginx  \ 
--user=www  \
--group=www \
--with-http_ssl_module  \
--with-http_stub_status_module \ 
--with-file-aio  \
--with-http_dav_module \
--with-pcre=/usr/local/src/pcre-8.37 && make && make install && chown -R www. /usr/local/nginx

COPY nginx.conf /usr/local/nginx/conf/nginx.conf
ADD vhosts /usr/local/nginx/conf/vhosts
ADD run.sh /root/run.sh
RUN chmod 755 /root/run.sh

CMD ["/root/run.sh"]

EXPOSE 80 443
