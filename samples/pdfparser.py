from PyPDF2 import PdfReader


reader = PdfReader("aws.pdf")
totalP = reader.pages
print(len(totalP))
page = reader.pages[24]
print(page.extract_text())

