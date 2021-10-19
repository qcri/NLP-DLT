const populateJSON = require("../../utils/data/populateJSON");
const invoke = require("../../utils/invoke/invoke");

const addOrg = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let method = "addOrg";
        let value = data.feature;
        let userDetails = data.userDetails;
        value = await invoke(method, value, userDetails.username);
        if (!value) {
            res.sendStatus(403);
        }
        await populateJSON(value)
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { addOrg };
