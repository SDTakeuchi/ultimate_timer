FROM  nginx:alpine

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./uwsgi_params /etc/nginx/uwsgi_params

CMD ["nginx", "-g", "daemon off;"]
