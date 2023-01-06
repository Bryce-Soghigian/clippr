import librosa
import numpy as np

def load_and_preprocess_music_data(data_dir, duration=30):
    # Load the audio files and extract features
    features = []
    labels = []
    file_names = os.listdir(data_dir)
    for file_name in file_names:
        file_path = os.path.join(data_dir, file_name)
        audio, sr = librosa.load(file_path, duration=duration)
        audio = audio.reshape(-1, 1)
        mfccs = librosa.feature.mfcc(audio, sr=sr)
        features.append(mfccs)
        labels.append(file_name.split(".")[0])
    features = np.array(features)
    labels = np.array(labels)
    
    # Normalize the features
    mean = np.mean(features, axis=0)
    std = np.std(features, axis=0)
    features = (features - mean) / std
    
    # Convert the features and labels to tensors
    features = torch.from_numpy(features).float()
    labels = torch.from_numpy(labels).long()
    
    # Create a dataset from the features and labels
    dataset = torch.utils.data.TensorDataset(features, labels)
    
    return dataset

