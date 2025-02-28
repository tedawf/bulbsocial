package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"math/rand"

	"github.com/tedawf/bulbsocial/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi",
	"ivan", "judy", "karl", "laura", "mallory", "nina", "oscar", "peggy",
	"quinn", "rachel", "steve", "trent", "ursula", "victor", "wendy", "xander",
	"yvonne", "zack", "amber", "brian", "carol", "doug", "eric", "fiona",
	"george", "hannah", "ian", "jessica", "kevin", "lisa", "mike", "natalie",
	"oliver", "peter", "queen", "ron", "susan", "tim", "uma", "vicky",
	"walter", "xenia", "yasmin", "zoe",
}

var titles = []string{
	"Buy Signal: Tesla (TSLA)",
	"Short Gold Futures",
	"EUR/USD Long Setup",
	"Buy Bitcoin (BTC)",
	"Sell Apple (AAPL)",
	"Short Oil Futures",
	"Buy Signal: Nvidia (NVDA)",
	"USD/JPY Short Setup",
	"Buy Signal: Gold",
	"Sell Netflix (NFLX)",
	"Long Natural Gas Futures",
	"Buy Tesla Weekly Options",
	"Sell EUR/GBP",
	"Buy Ethereum (ETH)",
	"Short Signal: Amazon (AMZN)",
	"Trade Alert: Silver Futures",
	"GBP/USD Buy Opportunity",
	"Crypto Bull Run Incoming?",
	"Sell Signal: Meta (META)",
	"Long-Term Play: Berkshire Hathaway",
	"Buy Signal: AMD",
	"Short Crude Oil Strategy",
	"High-Risk Play: Penny Stocks",
	"Forex Watch: AUD/CAD",
	"Day Trading Strategy: S&P Futures",
	"Buy Signal: Copper Futures",
	"Sell Opportunity: Google (GOOGL)",
	"Short Signal: Netflix (NFLX)",
	"Momentum Alert: Shopify (SHOP)",
	"Trade Setup: Tesla Calls",
	"Buy and Hold: Energy ETF",
	"Scalping Setup: NASDAQ Futures",
	"Market Dip Buy: Bitcoin",
	"Short Play: High-Yield Bonds",
	"Swing Trade: Disney (DIS)",
	"Buy Opportunity: Gold Miners ETF",
	"Earnings Alert: Microsoft (MSFT)",
	"Trend Reversal: EUR/JPY",
	"Crypto Breakout: Solana (SOL)",
	"Long Signal: Natural Gas ETF",
}

var contents = []string{
	"Buy TSLA at $200. Target price: $230. Stop loss: $190. Rationale: Strong earnings report and increased delivery guidance.",
	"Short Gold at $1950. Target price: $1900. Stop loss: $1970. Rationale: Hawkish Fed commentary and strengthening USD.",
	"Enter long EUR/USD at 1.0650. Target: 1.0800. Stop loss: 1.0600. Rationale: Positive EU economic data and weakening US job market.",
	"Buy BTC at $35,000. Target price: $40,000. Stop loss: $33,000. Rationale: Bullish chart pattern and institutional buying activity.",
	"Sell AAPL at $180. Target price: $170. Stop loss: $185. Rationale: Overvaluation concerns and declining iPhone sales.",
	"Short WTI Crude at $85. Target: $78. Stop loss: $87. Rationale: Increased US oil production and weak demand signals from China.",
	"Buy NVDA at $500. Target price: $550. Stop loss: $480. Rationale: AI demand continues to drive revenue growth.",
	"Enter short USD/JPY at 148.50. Target: 146.00. Stop loss: 149.50. Rationale: Intervention concerns from the Bank of Japan.",
	"Buy Gold at $1920. Target price: $1960. Stop loss: $1900. Rationale: Rising geopolitical tensions and dovish Fed signals.",
	"Sell NFLX at $420. Target price: $400. Stop loss: $430. Rationale: Weak subscriber growth in key regions.",
	"Go long on Natural Gas at $3.50. Target: $4.20. Stop loss: $3.20. Rationale: Anticipated colder winter and rising demand.",
	"Buy TSLA $210 weekly call options. Expiry: Nov 24, 2024. Premium: $5.50. Rationale: Anticipation of breakout post-earnings report.",
	"Sell EUR/GBP at 0.8700. Target: 0.8600. Stop loss: 0.8750. Rationale: Diverging interest rate outlooks for the ECB and BoE.",
	"Buy ETH at $2,000. Target price: $2,300. Stop loss: $1,900. Rationale: Bullish DeFi activity and staking demand.",
	"Short AMZN at $125. Target price: $115. Stop loss: $130. Rationale: Weak retail sales data and increasing competition.",
	"Trade Silver Futures at $23.50. Target price: $25.00. Stop loss: $22.75. Rationale: Increasing industrial demand.",
	"Buy GBP/USD at 1.2500. Target: 1.2650. Stop loss: 1.2400. Rationale: Strong UK economic data and weaker USD outlook.",
	"Crypto markets showing strong momentum. BTC and ETH expected to lead the next bull run. Focus on high liquidity pairs.",
	"Sell META at $350. Target price: $340. Stop loss: $360. Rationale: Overvaluation concerns and declining ad revenue.",
	"Long-term buy on BRK-A at $550,000. Target price: $580,000. Rationale: Strong fundamentals and diversified portfolio.",
	"Buy AMD at $90. Target price: $100. Stop loss: $85. Rationale: Increasing demand for semiconductor chips.",
	"Short Crude Oil at $80. Target price: $75. Stop loss: $82. Rationale: OPEC production increases and weakening global demand.",
	"High-risk buy on penny stocks in the renewable energy sector. Focus on emerging markets and innovative companies.",
	"Watch AUD/CAD for a breakout above 0.9200. Target price: 0.9350. Stop loss: 0.9100. Rationale: Positive Australian GDP data.",
	"Day trade on S&P futures at 4400. Target price: 4450. Stop loss: 4375. Rationale: Strong momentum post-FOMC meeting.",
	"Buy Copper Futures at $4.00. Target price: $4.50. Stop loss: $3.80. Rationale: Rising demand in renewable energy projects.",
	"Sell GOOGL at $135. Target price: $125. Stop loss: $140. Rationale: Regulatory concerns and slowing ad growth.",
	"Short NFLX at $420. Target price: $390. Stop loss: $430. Rationale: Missed subscriber growth expectations.",
	"Buy SHOP at $60. Target price: $70. Stop loss: $55. Rationale: Improved earnings and positive future guidance.",
	"Buy TSLA $250 weekly calls. Target premium: $15.00. Stop loss: $10.00. Rationale: Anticipated breakout post-earnings.",
	"Invest in Energy ETF XLE. Target price: $90. Stop loss: $80. Rationale: Rising energy prices and geopolitical tensions.",
	"Scalp NASDAQ futures at 15000. Target price: 15150. Stop loss: 14950. Rationale: Quick momentum-based opportunity.",
	"Buy BTC at $32,000. Target price: $38,000. Stop loss: $30,000. Rationale: Market rebound after recent correction.",
	"Short high-yield bonds ETF at $50. Target price: $45. Stop loss: $52. Rationale: Rising interest rate expectations.",
	"Swing trade DIS at $90. Target price: $100. Stop loss: $85. Rationale: Upcoming positive earnings catalyst.",
	"Buy Gold Miners ETF GDX at $30. Target price: $35. Stop loss: $28. Rationale: Gold price recovery expected.",
	"Monitor MSFT pre-earnings. Buy at $330. Target price: $360. Stop loss: $320. Rationale: Strong AI integration pipeline.",
	"Trend reversal expected on EUR/JPY at 150. Target price: 152. Stop loss: 148. Rationale: Policy divergence anticipation.",
	"Buy SOL at $20. Target price: $25. Stop loss: $18. Rationale: Increased network activity and NFT adoption.",
	"Long Natural Gas ETF at $10. Target price: $12. Stop loss: $9. Rationale: Anticipated colder winter boosting demand.",
}

var tags = [][]string{
	{"TSLA", "buy signal", "stock"},
	{"gold", "futures", "short"},
	{"EUR/USD", "forex", "long"},
	{"crypto", "BTC", "buy signal"},
	{"AAPL", "sell signal", "stock"},
	{"oil", "futures", "short"},
	{"NVDA", "AI", "buy signal"},
	{"USD/JPY", "forex", "short"},
	{"gold", "buy signal", "commodity"},
	{"NFLX", "sell signal", "stock"},
	{"natural gas", "futures", "long"},
	{"options", "TSLA", "weekly calls"},
	{"EUR/GBP", "forex", "sell signal"},
	{"ETH", "crypto", "buy signal"},
	{"AMZN", "short", "stock"},
	{"silver", "futures", "trade alert"},
	{"GBP/USD", "forex", "buy signal"},
	{"crypto", "BTC", "bull run"},
	{"META", "sell signal", "stock"},
	{"BRK-A", "long-term", "investment"},
	{"AMD", "buy signal", "semiconductors"},
	{"crude oil", "short", "strategy"},
	{"penny stocks", "renewables", "high risk"},
	{"AUD/CAD", "forex", "breakout"},
	{"S&P futures", "day trade", "strategy"},
	{"copper", "futures", "buy signal"},
	{"GOOGL", "sell signal", "regulatory concerns"},
	{"NFLX", "short signal", "stock"},
	{"SHOP", "momentum", "buy signal"},
	{"TSLA", "calls", "trade setup"},
	{"energy ETF", "buy and hold", "investment"},
	{"NASDAQ futures", "scalping", "momentum"},
	{"BTC", "buy signal", "crypto"},
	{"bonds", "short signal", "high yield"},
	{"DIS", "swing trade", "earnings"},
	{"gold miners", "ETF", "buy signal"},
	{"MSFT", "earnings", "AI"},
	{"EUR/JPY", "forex", "trend reversal"},
	{"SOL", "crypto", "breakout"},
	{"natural gas", "ETF", "long signal"},
}

var comments = []string{
	"Great post! Thanks for sharing.",
	"I completely agree with this analysis.",
	"This is really helpful. Appreciate the insight!",
	"Interesting perspective. I'll definitely consider this.",
	"Thanks for the update. Looking forward to more posts like this.",
	"Could you elaborate on this further?",
	"I wasn't aware of this. Thanks for pointing it out!",
	"This seems like a solid strategy.",
	"Very informative and well-written.",
	"I've been following this trend too. Good stuff!",
	"Wow, this really opened my eyes. Thanks!",
	"I'll keep this in mind for my next trade.",
	"Great breakdown of the details. Much appreciated.",
	"This post is gold. Thanks for the effort!",
	"Do you have any more insights on this?",
	"Totally agree with your point here.",
	"I'll definitely look into this. Thanks for sharing!",
	"Interesting take. I hadn't thought of it this way.",
	"This is exactly what I was looking for. Thanks!",
	"Thanks for simplifying this topic. It makes so much sense now.",
	"Great content as always. Keep it up!",
	"Very clear and concise. Well done.",
	"This is an excellent analysis. Thanks for the effort.",
	"Do you think this trend will continue?",
	"I really like this approach. Thanks for sharing!",
	"Solid advice. I'll try this out.",
	"This is quite insightful. Keep posting more!",
	"Thanks for the detailed explanation.",
	"I appreciate your perspective on this.",
	"This is really valuable information. Thanks a lot!",
	"Nice post! What are your thoughts on the long-term effects?",
	"Thanks for breaking this down. It makes a lot of sense.",
	"I'll definitely share this with others. Great insight!",
	"This helped me understand the topic much better.",
	"Thanks for keeping us updated!",
	"I've been following your posts, and this one is my favorite so far.",
	"This analysis aligns perfectly with my thoughts.",
	"Such an interesting viewpoint. Thanks for posting!",
	"This really helped clarify things for me.",
	"Awesome content. Keep it coming!",
	"Your posts are always on point. Thanks for sharing.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		randNum := rand.Intn(len(titles))

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[randNum],
			Content: contents[randNum],
			Tags:    tags[randNum],
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
