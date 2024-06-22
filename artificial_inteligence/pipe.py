from transformers import pipeline, AutoTokenizer, AutoModelForCausalLM  # type: ignore
from typing import List, Dict

tokenizer = AutoTokenizer.from_pretrained("microsoft/DialoGPT-large")
model = AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-large")

pipe: pipeline = pipeline("text-generation", tokenizer=tokenizer, model=model)

message = List[Dict[str, str]]
result = List[Dict[str, List[Dict[str, str]]]]


def get_chat(question: str) -> str:
    mes: message = [{"role": "user", "content": question}]
    res: result = pipe(mes)
    return res[0]["generated_text"][1]["content"]
