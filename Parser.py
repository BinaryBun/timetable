from selenium import webdriver
from time import *

def run_groups(list_of_groups, driver):
	for group in list_of_groups:
		driver.add_cookie({'name' : 'group', 'value' : group})
		driver.refresh()
		sleep(2)
		data = []
		for i in driver.find_elements_by_xpath('//div[@class="bold schedule-day__title"]'):
			data.append([i.text])
			for j in driver.find_elements_by_xpath(f'//div[text()="{data[-1][0]}"]/following-sibling::div/child::*/child::div[1]'):
				try:
					data[-1].append([j.text])
					req = f'//div[text()="{data[-1][0]}"]/following-sibling::div/child::*/child::div[1][text()="{j.text}"]/following-sibling::div/child::div[@class = "schedule-lesson  "]/div[@class != "clear"]'
					try:
						data[-1][-1].append(driver.find_element_by_xpath(req+'[1]/a').text)
					except:
						data[-1][-1].append(driver.find_element_by_xpath(req+'[1]').text)
					data[-1][-1].append(driver.find_element_by_xpath(req+'[2]').text)
					data[-1][-1].append(driver.find_element_by_xpath(req+'[3]/span').text)
					data[-1][-1].append(driver.find_element_by_xpath(req+'[4]').text)
				except:
					data[-1].pop()
		
		for i in data:
			for j in i:
				print(j)


def main():
	browser = webdriver.Chrome()
	browser.get("https://rasp.dmami.ru/")
	run_groups(['211-331'], browser)


if __name__ == '__main__':
	main()