# Using an official WordPress image as a parent image here
FROM wordpress:latest

# Setting up a custom PHP configuration
COPY php.ini /usr/local/etc/php/

# Setting up WP-CLI
RUN curl -O https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar \
    && chmod +x wp-cli.phar \
    && mv wp-cli.phar /usr/local/bin/wp

# Expose port 80 to allow external connections to the web server
EXPOSE 80

# Start WordPress
CMD ["apache2-foreground"]
