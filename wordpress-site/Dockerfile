# Use an official WordPress image as a parent image
FROM wordpress:latest

# Set up your custom PHP configuration
COPY php.ini /usr/local/etc/php/

# Set up WP-CLI
RUN curl -O https://raw.githubusercontent.com/wp-cli/builds/gh-pages/phar/wp-cli.phar \
    && chmod +x wp-cli.phar \
    && mv wp-cli.phar /usr/local/bin/wp

# Expose port 80 to allow external connections to the web server
EXPOSE 80

# Start WordPress
CMD ["apache2-foreground"]
