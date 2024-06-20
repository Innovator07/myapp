# Using the official WordPress image as a base here
FROM wordpress:latest

# Installing here necessary PHP extensions
RUN apt-get update && apt-get install -y \
    libfreetype6-dev \
    libjpeg62-turbo-dev \
    libpng-dev \
    libzip-dev \
    zip \
    unzip \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install gd zip

# Installing WP-CLI for managing WordPress from the command line here
RUN curl -O https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar \
    && chmod +x wp-cli.phar \
    && mv wp-cli.phar /usr/local/bin/wp

# Adding a custom PHP configuration
COPY ./php.ini /usr/local/etc/php/

# Installing and activating the Akismet plugin using WP-CLI
RUN wp plugin install akismet --activate --allow-root

# Setting up a default command to start the WordPress server
CMD ["apache2-foreground"]
