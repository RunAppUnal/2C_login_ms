FROM ruby:2.3

RUN mkdir /login-ms
WORKDIR /login-ms

ADD Gemfile /login-ms/Gemfile
ADD Gemfile.lock /login-ms/Gemfile.lock

RUN bundle install
ADD . /login-ms

EXPOSE 6004
