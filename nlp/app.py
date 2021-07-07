import json
from flask import Flask
from dotenv import load_dotenv
from library.Parsing.ParseToNLP import parseToNLP
from library.ManageJSON.UpdateFileV2 import updateFile
from library.Parsing.ParseToString import parseToString
from library.ManageEntities.DateFinder import dateFinder
from library.Parsing.ParseToAmzComph import parseToAmzCompreh
from library.ManagePDF.PdfToString import convert_pdf_to_string
from library.ManageEntities.RecovEntAndPhr import recoverPhrases
from library.ManageEntities.RecovEntAndPhr import recoverEntities
from library.ManageEntities.OrganizationFinder import organizationFinder

app = Flask(__name__)

text = convert_pdf_to_string('./input/Proximus_Direct_Wholesale_Roaming_access_Agreement--2020_08_01_2020-08-31-12-53-17_cache_.pdf')

#txtParsedToStr = parseToString(text)

txtParsedToNLP = parseToNLP(text)

readyToComprh = parseToAmzCompreh(txtParsedToNLP)

entitiesList = recoverEntities(readyToComprh)

phrasesList = recoverPhrases(readyToComprh)

# POPULATE DATE
date = dateFinder(entitiesList)
updateFile('./output/Roaming Agreements Output Template.json',"date",0,"hint", date)

# POPULATE ORGANIZATIONS
organizations = organizationFinder(entitiesList)
updateFile('./output/Roaming Agreements Output Template.json',"organization",1,"hint", organizations)

#   entity = stringFinder(txtParsedToStr, "Mainterms&conditionsBetween", ',')
#   updateFile('./output/Roaming Agreements Output Template.json',"operators", 0, "name", entity)

#   entity = stringFinder(txtParsedToStr, 'Hereinafterreferredtoas"', '"')
#   updateFile('./output/Roaming Agreements Output Template.json',"operators", 0, "alias", entity)

#   entity = stringFinder(txtParsedToStr, ')And', ',')
#   updateFile('./output/Roaming Agreements Output Template.json',"operators", 1, "name", entity)

#   entity = applyModel(txtParsedToNLP, 'CARDINAL', '/', 2)
#   updateFile('./output/Roaming Agreements Output Template.json',"agreement_date", 0, "0", entity)

print(a)

@app.route('/')
def hello_world():
    return