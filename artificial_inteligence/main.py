from transformers import pipeline # type: ignore
from typing import List

messages: List[dict[str, str]] = [
    {"role": "user", "content": "Who are you?"},
]
pipe: pipeline = pipeline("text-generation", model="alpindale/magnum-72b-v1")
result: str = pipe(messages)

print(result)