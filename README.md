# Alexa Skill - Dice Roll

You're playing a dice game, but instead of bringing all the dice, you brought a device with Amazon Alexa on it. Strange, but ok. Fortunately, this Alexa Skill will help you out.

* "Alexa, ask Dice Roll 4 d 6" ... "Rolling 4 d 6 ... 13"
* "Alexa, ask Dice Roll d 20" ... "Rolling 1 d 20 ... 20"

## Deploy to AWS Lambda

The backend for this Alexa Skill is an AWS Lambda function, written in Golang, deployed via [Apex](https://apex.run)

```
apex deploy
```

Read https://www.starkandwayne.com/blog/quick-guide-to-deploying-golang-to-aws-lambda-using-apex/ for an introduction to deploying Golang to AWS Lambda using Apex.

Once the Lambda function is deployed, you'll then need to add an Alexa Skills Kit trigger to it.

![lambda-trigger](https://cl.ly/3f1B3l1J1b3w/download/Image%202017-01-10%20at%203.34.50%20PM.png)

Then, in your [Alexa Skills List](https://developer.amazon.com/edw/home.html#/skills/list), click "Add a New Skill", and set it up.

The Intent Schema will be:

```json
{
  "intents": [
    {
      "intent": "RollDiceIntent",
      "slots": [
        {
          "name": "HowMany",
          "type": "AMAZON.NUMBER"
        },
        {
          "name": "DiceSides",
          "type": "AMAZON.NUMBER"
        }
      ]
    }
  ]
}
```

The Sample Utterances will be:

```
RollDiceIntent {HowMany} d {DiceSides}
```
