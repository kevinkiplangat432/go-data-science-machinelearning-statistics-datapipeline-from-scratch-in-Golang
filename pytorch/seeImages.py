#!/usr/bin/env python
import torch
from torchvision import datasets
from torchvision.transforms import ToTensor
import matplotlib.pyplot as plt

# Same dataset from the quickstart
training_data = datasets.FashionMNIST(
    root="data",
    train=True,
    download=True,
    transform=ToTensor()
)

labels_map = {
    0: "T-Shirt", 1: "Trouser", 2: "Pullover", 3: "Dress", 4: "Coat",
    5: "Sandal", 6: "Shirt", 7: "Sneaker", 8: "Bag", 9: "Ankle Boot",
}

# Show a 4x4 grid of random samples
figure, axes = plt.subplots(4, 4, figsize=(8, 8))
figure.suptitle("FashionMNIST Samples", fontsize=14)

for ax in axes.flat:
    idx = torch.randint(len(training_data), size=(1,)).item()
    img, label = training_data[idx]
    
    ax.imshow(img.squeeze(), cmap="gray")
    ax.set_title(labels_map[label], fontsize=9)
    ax.axis("off")

plt.tight_layout()
plt.show()