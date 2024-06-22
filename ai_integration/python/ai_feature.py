import pickle
import numpy as np

# Load pre-trained model
with open('model.pkl', 'rb') as file:
    model = pickle.load(file)

def predict(data):
    data = np.array(data).reshape(1, -1)
    prediction = model.predict(data)
    return prediction

if __name__ == '__main__':
    sample_data = [1.0, 2.0, 3.0]
    prediction = predict(sample_data)
    print(f"Prediction: {prediction}")
