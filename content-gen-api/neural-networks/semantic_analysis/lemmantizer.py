import nltk
from textblob import TextBlob

# Download the NLTK data and WordNet lemmatizer
nltk.download("averaged_perceptron_tagger")
nltk.download("wordnet")
lemmatizer = nltk.WordNetLemmatizer()
def analyze_emotion(lyrics: str) -> str:
    """
    Analyze the emotion of the lyrics.
    
    Parameters:
    lyrics (str): The lyrics of the song.
    
    Returns:
        int: Sentiment score.
    """
    tokens = nltk.word_tokenize(lyrics)
    
    pos_tags = nltk.pos_tag(tokens)
    adj_tags = [tag[0] for tag in pos_tags if tag[1] == "JJ"]
    adj_lemmas = [lemmatizer.lemmatize(adj) for adj in adj_tags]
    
    # Use TextBlob to determine the sentiment of the lemmatized adjectives
    sentiment = 0.0
    for adj in adj_lemmas:
        sentiment += TextBlob(adj).sentiment.polarity
    
   return sentiment


