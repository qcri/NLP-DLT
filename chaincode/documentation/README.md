# Chaincode documentation
- This section is divided in `chaincode design`, `chaincode implementation` and `chaincode testing`.

## Chaincode design
1. `cd ~/NLP-DLT/chaincode/design`
2. To open the file `diagram_sequence_chaincode_v5.drawio` using [App Diagrams Tool](https://app.diagrams.net/)
3. To open the file `class_diagram_chaincode_v5.drawio` using [App Diagrams Tool](https://app.diagrams.net/)
4. To open the file `states_diagram.drawio` using [App Diagrams Tool](https://app.diagrams.net/)

### Register organization
- Any MNO must be registered before drafting a roaming agreement.
- Identity is verified at each interaction.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/registerOrg1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/registerOrg2.png">

### Start Agreement
- A registered organization can enable the drafting of a Roaming Agreement.
- Identity is verified.
- The inputs are two `json org` and `json ROAT.json`, i.e., the output of the NLP engine.
- The `RAID` is generated.
    - `RAID` is accesible for all MNOs.
- An event is emitted to set the state `ra_started`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/startAgreement1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/startAgreement2.png">

### Start Agreement Confirmation
- For the roaming agreement drafting to be valid, the other MNO must confirm it.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_ra_started`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmStartAgreement1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmStartAgreement2.png">

### Articles Update
- The drafting of the Roaming Agreement involves the updating of the articles. 
- Identity is verified.
- The inputs are `json org`, `RAID`, `article_num` and `jsonArticle`.
- An event is emitted to set the state `proposed_change`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/setArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/setArticle2.png">

### Confirmation of Article Update
- The other MNO must validate the article update.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_proposed_change`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmSetArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmSetArticle2.png">

### Article Deletion
- The drafting of the Roaming Agreement involves the deletion of the articles. 
- Identity is verified.
- The inputs are `json org`, `RAID` and `article_num`.
- An event is emitted to set the state `proposed_deletion`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/deleteArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/deleteArticle2.png">

### Confirmation of Article Deletion
- The other MNO must validate the article deletion.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_proposed_deletion`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmDeleteArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmDeleteArticle2.png">

### Agreement Achieved
- The drafting of the Roaming Agreement involves the acceptation of the drafting process. 
- Identity is verified.
- The inputs are `json org`and `RAID`.
- An event is emitted to set the state `RA_accepted`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/agreementAchieved1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/agreementAchieved2.png">

### Confirmation of Agreement Achieved
- The other MNO must validate the article deletion.
- Identity is verified.
- The inputs are `json org` and `RAID`.
- An event is emitted to set the state `confirm_RA_achieved`.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmAgreementAchieved1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/confirmAgreementAchieved2.png">

### Query Single Article
- Identity is verified.
- The inputs are `json org`, `RAID`and `article_num`.
- The content of `article_num` is returned.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/querySingleArticle1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/querySingleArticle2.png">

### Query All Article
- Identity is verified.
- The inputs are `json org` and `RAID`.
- The content of `jsonRA` is returned.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/queryAllArticles1.png">       
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/queryAllArticles2.png">

### State-to-state-transition
- Actions implies change of state. 
- The chaincode validates the changes of states.
<img src="https://github.com/sfl0r3nz05/NLP-DLT/blob/main/images/states_diagram.png">

### States and events
- The following table associates states to events emitted:

| States        | Events        |
| ------------- |:-------------:|
|               | org_created   |
| ra_started    | ra_started    |
| confirm_ra_started | confirm_ra_started      |
| proposed_change | proposed_change      |
| confirm_proposed_change | confirm_proposed_change      |
| ra_accepted | ra_accepted      |
| confirm_ra_achieved | confirm_ra_achieved      |