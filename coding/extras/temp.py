from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
from bs4 import BeautifulSoup
from webdriver_manager.chrome import ChromeDriverManager

import time
import csv

# Setup Chrome options
options = Options()
options.add_argument('--headless')  # Run in headless mode
options.add_argument('--disable-gpu')
options.add_argument('--no-sandbox')

# Set path to chromedriver

driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()), options=options)


# Target URL
url = "https://www.quikr.com/jobs/mumbai-vacancy+zwqxj2726005330"
driver.get(url)

# Scroll until no more content loads
SCROLL_PAUSE_TIME = 2
last_height = driver.execute_script("return document.documentElement.scrollHeight")

while True:
    driver.execute_script("""
        var scrollingElement = document.scrollingElement || document.documentElement;
        scrollingElement.scrollTop = scrollingElement.scrollHeight;
    """)
    time.sleep(SCROLL_PAUSE_TIME)
    new_height = driver.execute_script("return document.documentElement.scrollHeight")
    if new_height == last_height:
        break
    last_height = new_height

# After scrolling, parse the HTML
soup = BeautifulSoup(driver.page_source, 'html.parser')
driver.quit()

# Find job cards
job_cards = soup.find_all('div', class_='job-card')

# Extract job info
job_data = []
for card in job_cards:
    title_elem = card.find('a', class_='job-title')
    company_elem = card.find('div', class_='attributeVal cursor-default')
    salary_elem = card.find('div', class_='perposelSalary')
    location_elem = card.find('span', class_='city')

    title = title_elem.text.strip() if title_elem else ''
    link = title_elem['href'] if title_elem else ''
    company = company_elem.text.strip() if company_elem else ''
    salary = salary_elem.text.strip() if salary_elem else ''
    location = location_elem.text.strip() if location_elem else ''

    job_data.append([title, company, salary, location, link])

# Save to CSV
with open('quikr_jobs.csv', 'w', newline='', encoding='utf-8') as file:
    writer = csv.writer(file)
    writer.writerow(['Title', 'Company', 'Salary', 'Location', 'Link'])
    writer.writerows(job_data)

print(f"✅ Scraped {len(job_data)} jobs and saved to 'quikr_jobs.csv'")
