from selenium import webdriver
from time import *
import db

def set_map(audit):
	audit = audit.replace(" ", "")[:3]
	data = {"ав1": "https://yandex.ru/maps/213/moscow/?ll=37.646117%2C55.704199&mode=poi&poi%5Bpoint%5D=37.645163%2C55.704191&poi%5Buri%5D=ymapsbm1%3A%2F%2Forg%3Foid%3D1701016484&utm_source=main_stripe_big&z=20",
			"ав2": "https://yandex.ru/maps/213/moscow/?ll=37.645775%2C55.704509&mode=poi&poi%5Bpoint%5D=37.645704%2C55.704561&poi%5Buri%5D=ymapsbm1%3A%2F%2Forg%3Foid%3D167848998809&utm_source=main_stripe_big&z=20",
			"ав3": "https://yandex.ru/maps/213/moscow/?ll=37.645370%2C55.704667&mode=whatshere&utm_source=main_stripe_big&whatshere%5Bpoint%5D=37.646880%2C55.704792&whatshere%5Bzoom%5D=19&z=20",
			"ав4": "https://yandex.ru/maps/213/moscow/?ll=37.645370%2C55.704667&mode=poi&poi%5Bpoint%5D=37.646639%2C55.704452&poi%5Buri%5D=ymapsbm1%3A%2F%2Forg%3Foid%3D1737978898&utm_source=main_stripe_big&z=20",
			"ав5": "https://yandex.ru/maps/213/moscow/?ll=37.644590%2C55.705085&mode=poi&poi%5Bpoint%5D=37.646804%2C55.705504&poi%5Buri%5D=ymapsbm1%3A%2F%2Forg%3Foid%3D47532892840&utm_source=main_stripe_big&z=20",
			"ав6": "https://yandex.ru/maps/213/moscow/?ll=37.644590%2C55.705085&mode=poi&poi%5Bpoint%5D=37.646083%2C55.704282&poi%5Buri%5D=ymapsbm1%3A%2F%2Forg%3Foid%3D178922065429&utm_source=main_stripe_big&z=20",
			"Пр2": "https://yandex.ru/maps/213/moscow/?ll=37.543975%2C55.833983&mode=poi&poi%5Bpoint%5D=37.543473%2C55.833741&poi%5Buri%5D=ymapsbm1%3A%2F%2Forg%3Foid%3D126963524657&utm_source=main_stripe_big&z=20",
			"Пр1", "https://yandex.ru/maps/213/moscow/?ll=37.544323%2C55.833470&mode=poi&poi%5Buri%5D=ymapsbm1%3A%2F%2Forg%3Foid%3D1917526617&utm_source=main_stripe_big&z=20"}

	if audit in data:
		return data[audit]

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

		data_base = db.DB()
		for i in data:
			data_base.insert(i)

def main():
	browser = webdriver.Chrome()
	browser.get("https://rasp.dmami.ru/")
	run_groups(['211-331'], browser)

if __name__ == '__main__':
	main()
