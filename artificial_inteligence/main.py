# Use a pipeline as a high-level helper
from transformers import pipeline, AutoTokenizer, AutoModelForCausalLM  # type: ignore
from typing import Dict, List

tokenizer = AutoTokenizer.from_pretrained("microsoft/DialoGPT-large")
model = AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-large")

pipe: pipeline = pipeline("text-generation", tokenizer=tokenizer, model=model)

message = Dict[str, str]

messages: List[message] = [
    {
        "role": "user",
        "content": "who are you?",
    },
]

result: List[Dict[str, List[Dict[str, str]]]] = pipe(messages)


generated_texts = result[0]["generated_text"]

print(result)
for item in generated_texts:
    print(item["content"])
