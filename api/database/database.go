package database

import "api/models"

var questions = []models.Question{
    {
        ID:       1,
        Question: "What does the term 'bull market' mean?",
        Options: []models.Option{
            {ID: 1, Text: "A market where prices are falling"},
            {ID: 2, Text: "A market where prices are rising"},
            {ID: 3, Text: "A market with no significant price movement"},
            {ID: 4, Text: "A market that trades primarily in livestock"},
        },
        CorrectAnswer: 2,
    },
    {
        ID:       2,
        Question: "Which of the following is a common risk management strategy in trading?",
        Options: []models.Option{
            {ID: 1, Text: "Ignoring market trends"},
            {ID: 2, Text: "Using stop-loss orders"},
            {ID: 3, Text: "Investing all capital in one stock"},
            {ID: 4, Text: "Relying on rumors"},
        },
        CorrectAnswer: 2,
    },
    {
        ID:       3,
        Question: "What is 'short selling'?",
        Options: []models.Option{
            {ID: 1, Text: "Buying shares with the expectation that prices will rise"},
            {ID: 2, Text: "Selling shares you do not own, expecting the price to fall"},
            {ID: 3, Text: "Selling shares for a short period of time"},
            {ID: 4, Text: "Buying and selling shares within a single day"},
        },
        CorrectAnswer: 2,
    },
    {
        ID:       4,
        Question: "What is a 'blue chip' stock?",
        Options: []models.Option{
            {ID: 1, Text: "A stock that has a blue-colored stock certificate"},
            {ID: 2, Text: "A stock with a high risk and high reward"},
            {ID: 3, Text: "A stock from a large, well-established, and financially sound company"},
            {ID: 4, Text: "A stock that is performing poorly"},
        },
        CorrectAnswer: 3,
    },
    {
        ID:       5,
        Question: "Which economic indicator is often considered when trading Forex?",
        Options: []models.Option{
            {ID: 1, Text: "Gross Domestic Product (GDP)"},
            {ID: 2, Text: "Consumer Price Index (CPI)"},
            {ID: 3, Text: "Unemployment rate"},
            {ID: 4, Text: "All of the above"},
        },
        CorrectAnswer: 4,
    },
    {
        ID:       6,
        Question: "What does 'P/E ratio' stand for in trading?",
        Options: []models.Option{
            {ID: 1, Text: "Price/Earnings ratio"},
            {ID: 2, Text: "Profit/Equity ratio"},
            {ID: 3, Text: "Price/Equity ratio"},
            {ID: 4, Text: "Profit/Earnings ratio"},
        },
        CorrectAnswer: 1,
    },
    {
        ID:       7,
        Question: "Which of the following is a key principle of technical analysis?",
        Options: []models.Option{
            {ID: 1, Text: "Analyzing company earnings reports"},
            {ID: 2, Text: "Studying past price movements and trading volumes"},
            {ID: 3, Text: "Evaluating macroeconomic indicators"},
            {ID: 4, Text: "Consulting with a financial advisor"},
        },
        CorrectAnswer: 2,
    },
    {
        ID:       8,
        Question: "What is 'diversification' in trading?",
        Options: []models.Option{
            {ID: 1, Text: "Investing all your money in one type of asset"},
            {ID: 2, Text: "Spreading investments across various assets to reduce risk"},
            {ID: 3, Text: "Buying and selling the same stock multiple times in a day"},
            {ID: 4, Text: "Consolidating all investments into a single portfolio"},
        },
        CorrectAnswer: 2,
    },
    {
        ID:       9,
        Question: "Which type of order is used to limit potential losses on a trade?",
        Options: []models.Option{
            {ID: 1, Text: "Market order"},
            {ID: 2, Text: "Stop-loss order"},
            {ID: 3, Text: "Limit order"},
            {ID: 4, Text: "Good-till-canceled order"},
        },
        CorrectAnswer: 2,
    },
    {
        ID:       10,
        Question: "What does the term 'liquidity' refer to in trading?",
        Options: []models.Option{
            {ID: 1, Text: "The amount of cash available in a trader's account"},
            {ID: 2, Text: "The ease with which an asset can be bought or sold in the market without affecting its price"},
            {ID: 3, Text: "The amount of leverage a trader uses"},
            {ID: 4, Text: "The stability of a financial market"},
        },
        CorrectAnswer: 2,
    },
}

func GetQuestions() []models.Question {
    return questions
}

func CheckAnswers(userAnswers []models.UserAnswer) models.QuizResult {
    var result models.QuizResult

    for _, userAnswer := range userAnswers {
        for _, question := range questions {
            if question.ID == userAnswer.QuestionID {
                if question.CorrectAnswer == userAnswer.AnswerID {
                    result.CorrectAnswers++
                } else {
                    result.IncorrectAnswers++
                }
                break
            }
        }
    }

    return result
}
