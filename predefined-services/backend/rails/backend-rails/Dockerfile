FROM ruby:${{RAILS_RUBY_VERSION}}
WORKDIR /usr/src/app

RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
RUN apt-get update && apt-get install -y --no-install-recommends nodejs yarn && rm -rf /var/lib/apt/lists/*
RUN yarn -v
#? if services.database contains name == "postgres" {
#RUN apt-get install -y --no-install-recommends postgresql-client && rm -rf /var/lib/apt/lists/*  
#? }

COPY Gemfile* ./
RUN bundle install
COPY . .
RUN bundle install
RUN bundle exec rails webpacker:install

EXPOSE 3000
CMD ["rails", "server", "-b", "0.0.0.0"]