from time import sleep
from selenium.webdriver.common.by import By
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

input_file = "yt_ids.txt"
downsub_url = 'https://downsub.com/'

def download_srt(url):
    driver = webdriver.Safari()
    driver.get(downsub_url)
    # search
    search_field = driver.find_element(By.CSS_SELECTOR, "input[name='url']")
    search_field.send_keys(url)
    search_field.submit()
    sleep(5)
    # it takes a moment for the download and translation to complete
    selector = "//div[@class='layout justify-start align-center']/button"
    buttons = driver.find_elements(By.XPATH, selector)
    if not len(buttons):
        driver.quit()
        raise TimeoutError
    srt_download = buttons[0]
    # srt_download = WebDriverWait(driver, 6).until(
    #     EC.presence_of_element_located((By.XPATH, selector))
    # )
    # click download
    srt_download.click()
    # get youtube title from html
    text = driver.find_element(By.XPATH, "//div[@id='ds-information']/div/a").text
    print("video title:", text)
    sleep(7) # allow download to start
    driver.quit()

if __name__ == "__main__":
    with open(input_file) as file:
        for line in file:
            new_url = "https://youtu.be/" + line.strip()
            try:
                download_srt(new_url)
            except TimeoutError:
                print("ERR: Failed to download:", line.strip())
                continue
            print("downloaded srt for:", new_url)