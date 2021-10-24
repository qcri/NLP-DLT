const populateLER = require("../../utils/data/populateLER");
const invokeEvents = require("../../utils/invoke/invokeEvents");

const addOrg = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "addOrg";
        let value = data.feature;
        let userDetails = data.userDetails;
        if (!userDetails || !value.mno_name || !value.mno_country || !value.mno_network) {
            res.sendStatus(201);
            res.end("201");
            return
        }
        let noArg = 1
        let event_name = "created_org"
        value, eventHf = await invokeEvents(method, event_name, noArg, value, "", "", userDetails.username);
        if (!value) {
            res.sendStatus(403);
        }
        await populateLER(value.mno_name)
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { addOrg };
