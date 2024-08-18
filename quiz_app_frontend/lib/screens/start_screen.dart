import 'package:flutter/material.dart';

class StartScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        decoration: BoxDecoration(
          gradient: LinearGradient(
            colors: [
              Color(0xFF5A4FCF),
              Color(0xFF6A1B9A),
            ],
            begin: Alignment.topLeft,
            end: Alignment.bottomRight,
          ),
        ),
        child: SafeArea(
          child: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(
                      'Quiz',
                      style: TextStyle(
                        color: Colors.white,
                        fontSize: 28,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    Row(
                      children: [
                        // Commented out the crown icon
                        // FaIcon(FontAwesomeIcons.crown, color: Colors.white),
                        SizedBox(width: 5),
                        Text(
                          '20',
                          style: TextStyle(color: Colors.white),
                        ),
                      ],
                    ),
                  ],
                ),
                SizedBox(height: 20),
                Text(
                  'Popular Game 🔥',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 22,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                SizedBox(height: 10),
                // Swipeable horizontal list of quiz cards
                Expanded(
                  child: ListView(
                    scrollDirection: Axis.horizontal,
                    children: [
                      _buildQuizCard(
                        context,
                        title: 'Technology',
                        description: 'Explore the world of technology...',
                        imageAsset: 'assets/quiz.png',
                      ),
                      _buildQuizCard(
                        context,
                        title: 'Social Sciences',
                        description: 'Dive into social sciences...',
                        imageAsset: 'assets/quiz.png',
                      ),
                      _buildQuizCard(
                        context,
                        title: 'Science',
                        description: 'Unravel the mysteries of science...',
                        imageAsset: 'assets/quiz.png',
                      ),
                    ],
                  ),
                ),
                SizedBox(height: 20),
                Text(
                  'Recent Played',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 22,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                SizedBox(height: 5),
                // Recent played list with a single card (can be repeated as needed)
                _buildQuizCard(
                  context,
                  title: 'Social Sciences',
                  description: 'Played recently...',
                  imageAsset: 'assets/quiz.png',
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildQuizCard(BuildContext context,
      {required String title,
      required String description,
      required String imageAsset}) {
    return Container(
      width: 160, // Smaller width for more compact swiping
      margin: EdgeInsets.only(right: 15), // Adds space between cards
      padding: EdgeInsets.all(12),
      decoration: BoxDecoration(
        color: Colors.orangeAccent,
        borderRadius: BorderRadius.circular(20),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Image.asset(
            imageAsset,
            height: 70, // Reduced height to fit content
            width: 140,
            fit: BoxFit.cover,
          ),
          SizedBox(height: 5),
          Text(
            title,
            style: TextStyle(
              color: Colors.white,
              fontSize: 16,
              fontWeight: FontWeight.bold,
            ),
          ),
          SizedBox(height: 5),
          Text(
            description,
            style: TextStyle(
              color: Colors.white70,
              fontSize: 12,
            ),
          ),
          SizedBox(height: 5),
          ElevatedButton(
            onPressed: () {
              // Navigate to the quiz page using the correct route name
              Navigator.pushNamed(context, '/quiz');
            },
            style: ElevatedButton.styleFrom(
              backgroundColor:
                  Colors.white, // Use backgroundColor instead of primary
              foregroundColor: Colors
                  .orangeAccent, // Use foregroundColor instead of onPrimary
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(20),
              ),
              padding: EdgeInsets.symmetric(horizontal: 10, vertical: 8),
            ),
            child: Text(
              'Start Quiz',
              style: TextStyle(
                fontSize: 14,
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
        ],
      ),
    );
  }
}
