# Amazon Product Scraper Assignment

This code is the solution for Seller App Assignment


# How to Use:
open a terminal. Now run
- `git clone https://github.com/Ayush-Walia/amazon-scraper`
- `cd amazon-scraper`
- `docker compose -f docker-compose-prod.yml up`

**Note**: You can use pass ```docker-compose.yml``` in docker-compose command if you want to run the service from code instead.

These commands should spin up three containers with services `scraper-service`, `scrape-data-manager` and `mongodb`.
- **scraper-service** is listening on **http://localhost:5000**
- **scrape-data-manager** is listening on **http://localhost:5001**
- **mongodb** is listening on **http://localhost:27017**

Once the containers are up, you can hit the following API endpoints:
1. `http://localhost:5000/healthcheck` (Allowed Method: GET) - This is the Health-check API for scraper-service
2. `http://localhost:5001/healthcheck` (Allowed Method: GET) - This is the Health-check API for scrape-data-manager
3. `http://localhost:5000/scrape_page` (Allowed Method: POST) - This is the API for scraping a page and storing the data in mongodb

   API Payload Structure:

    ```
    Content-Type: application/json
    Payload Body: {"url": Amozon url which you want to hit (string) }
    ```

   Example Curl Request:
    ```(JSON)
    curl --location --request POST 'localhost:5000/scrape_page' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "url": "https://www.amazon.com/Samsung-Unlocked-Smartphone-Pro-Grade-SM-G991UZVAXAA/dp/B08N3J7GJ8/ref=sr_1_3?crid=2HADJ96HHA7PG&dchild=1&keywords=samsung+galaxy+s21&qid=1633236578&sprefix=samsung+%2Caps%2C422&sr=8-3"
    }'
    ```

4. `http://localhost:5001/save_product` (Allowed Method: POST) - This is the API for save the data in Database. It is called by scraper-service

    API Payload Structure:

     ```
     Content-Type: application/json
     Payload Body: {
                 "url": Amozon url which you want to hit (string),
                 "product": {
                     "name": Name of the product (string),
                     "image_url": URL of product image (string),
                     "description": List containing the description ([]string),
                     "price": Price (string),
                     "total_reviews": Number of reviews (integer)
                 }
     }
     ```

     Example Curl Request:

     ```(JSON)
     curl --location --request POST 'http://localhost:5001/save_product' \
     --header 'Content-Type: application/json' \
     --data-raw '{
        "url": "https://www.amazon.com/Samsung-Unlocked-Smartphone-Pro-Grade-SM-G991UZVAXAA/dp/B08N3J7GJ8/ref=sr_1_3?crid=2HADJ96HHA7PG&dchild=1&keywords=samsung+galaxy+s21&qid=1633236578&sprefix=samsung+%2Caps%2C422&sr=8-3",
        "product": {
        "name": "Samsung Galaxy S21 5G | Factory Unlocked Android Cell Phone | US Version 5G Smartphone | Pro-Grade Camera, 8K Video, 64MP High Res | 128GB, Phantom Violet (SM-G991UZVAXAA)",
        "image_url": "https://m.media-amazon.com/images/I/512Va+-kCJL._AC_SY300_SX300_.jpg",
        "description": [
        "Pro Grade Camera: Zoom in close, take photos and videos like a pro, and capture incredible share-ready moments with our easy-to-use, multi-lens camera",
        "Sharp 8K Video: Capture your life’s best moments in head-turning, super-smooth 8K video that gives your movies that cinema-style quality",
        "Multiple Ways to Record: Create share-ready videos and GIFs on the spot with multi-cam recording and automatic professional-style effects",
        "30 Space Zoom: Get amazing power and clarity, zoom in from afar or magnify details of nearby objects; Zoom Lock keeps focus and stability",
        "Higher Resolution: 64 MP camera system captures and shares detailed portraits, stunning landscapes and crisp close-ups",
        "All Day Intelligent Battery: Intuitively manages your cellphone’s usage to conserve energy, so you can go all day without charging (based on average battery life under typical usage conditions)",
        "Power of 5G: Get next-level power for everything you love to do with Galaxy 5G; More sharing, more gaming, more experiences and never miss a beat\n\nWireless communication technology: CellularCamera description: FrontThe Galaxy S21 5G does not currently support eSim in the U.S."
        ],
        "price": "$799.99",
        "total_reviews": 1355
        }
     }'
     ```

You can verify the data in mongodb by using any mongo client. Connect to localhost: 27017 and no authentication is required.