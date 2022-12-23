import nltk
from string import punctuation
from termcolor import colored

students_akadem = {
    "Щербина": 1,
    "Кузнецов": 1,
    "Кузнецова": 1,
    "Иванов": 5,
    "Петров": 3,
    "Сидоров": 2,
    "Сидорова": 8,
    "Алексеев": 5,
    "Романов": 3,
    "Романова": 3,
    "Сергеев": 2,
    "Сергеева": 5,
    "Александров": 2,
    "Александрова": 9,
    "Андреев": 1,
}

# эксп часть
# фамилии в принимающие в участие
# график
# Примеры запросов
# обязательно вопрос не распознается

def expert_opinion(feature):
    if feature == 'много':
        return 8
    elif feature == 'средне':
        return 5
    elif feature == 'мало':
        return 2
    elif feature == 'немало':
        return 7
    elif feature == 'немного':
        return 4
    elif feature == 'оченьнемало':
        return 8
    elif feature == 'оченьмного':
        return 9
    elif feature == 'оченьмало':
        return 0
    elif feature == 'оченьнемного':
        return 1
    elif feature == 'неоченьмного':
        return 3
    elif feature == 'неоченьмало':
        return 7
    return None

def replace_merged_words(toks):
    new_toks = []
    for i in range(len(toks)):
        if toks[i] == "немного":
            new_toks.append("не")
            new_toks.append("много")
        elif toks[i] == "немало":
            new_toks.append("не")
            new_toks.append("мало")
        else:
            new_toks.append(toks[i])
    return new_toks


def mb_before_word(tokens, feat_index, offset=-1):
    if feat_index + offset < 0:
        return ""
    if tokens[feat_index + offset] not in ["не", "очень"]:
        return ""
    return tokens[feat_index + offset]

def perror(*msg):
    if not isinstance(msg, str):
        msg = " ".join([str(m) for m in msg if not isinstance(m, str) or m])
    print(colored(msg, 'red'))

text = "Кто немного раз ходил в академ?"

tokens = nltk.word_tokenize(text.lower(), language="russian")
tokens = [token for token in tokens if token not in punctuation]
tokens = replace_merged_words(tokens)

# В вопросе спрашивают про академ?
print('Разбитые токены:', tokens)

if not 'академ' in tokens:
    perror("В вопросе нет слова 'академ'")
    exit(1)

# определяем базовый признак
# мало, средне, много
base_words = ['много', 'мало', 'средне']
base_words_count = {i: 0 for i in base_words}

# Считаем количество вхождений слов из base_words
# в список токенов
for token in tokens:
    if token in base_words:
        base_words_count[token] += 1

if sum(base_words_count.values()) == 0:
    perror('В вопросе нет слов из списка', base_words)
    exit(1)

# Если в вопросе есть одно вхождение слова из base_words
# то это признак

found_base_word = None
for word, count in base_words_count.items():
    if count == 1:
        found_base_word = word
        break

for other_word, count in base_words_count.items():
    if count > 0 and other_word != found_base_word:
        perror('В вопросе больше одного вхождения слова', other_word, 'и слова', found_base_word)
        exit(1)

existing_features = [i for i in base_words_count if base_words_count[i] != 0]
feature = existing_features[0]
feature_index_in_sentence = tokens.index(feature)


feature = mb_before_word(tokens, feature_index_in_sentence, -2) + mb_before_word(tokens, feature_index_in_sentence) + feature

print("Количество уходов в академ")
print('Признак:', feature)
print('Оценка:', expert_opinion(feature))

print("Студенты которые ушли в академ ", feature, f", ({expert_opinion(feature)}) раз")
for student, akadem in students_akadem.items():
    if akadem == expert_opinion(feature):
        print(student)