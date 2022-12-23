import nltk
# nltk.download('punkt')
from string import punctuation
from termcolor import colored

conties_murav = {
    'Балтика 3': 3,
    'Балтика 1': 1,
    'Балтика 2': 2,
    'Балтика 6': 5,
    'Мартини': 4,
    'Балтика 5': 3,
    'Старый мельник': 2,
    'Клинское': 3,
    'Деревенька': 5,
    'Крушовиче': 3,
    'Апсент': 5,
    'Мотор': 3,
    'Аррарат': 4,
    'Озерная': 5,
    'Нету меда': 1
}


def expert_opinion(feature):
    if feature == 'легкий':
        return 1
    elif feature == 'средний':
        return 2
    elif feature == 'интеренсный':
        return 3
    elif feature == 'высокий':
        return 4
    elif feature == 'не употребительный':
        return 5
    return None

def mass_to_feature(mass):
    ret = []
    prov = []
    for i in mass:
        if i not in prov:
            prov.append(i)
    for i in prov:
        
        if i == 1:
            ret.append('легкий')
        if i == 2:
            ret.append('средний')
        if i == 3:
            ret.append('интеренсный')
        if i == 4:
            ret.append('высокий')
        if i == 5:
            ret.append('не употребительный')
            
    return ret

def perror(*msg):
    if not isinstance(msg, str):
        msg = " ".join([str(m) for m in msg if not isinstance(m, str) or m])
    print(colored(msg, 'red'))

def take_murav(feature):
    slov = 'спирт'
    for i in range(len(slov)):
        if slov[i] != feature[i]:
            return False
    return True

def take_murav_from_tokens(tokens):
    for i in tokens:
        if take_murav(i):
            return True
    return False


def find_terms(tokens):
    ret = []
    flag = 0
    for token in tokens:
        if flag == 5    :
            if token == 'употребительный':
                ret.append(5)
        if token == 'легкий':
            ret.append(1)
        elif token == 'средний':
            ret.append(2)
        elif token == 'интеренсный':
            ret.append(3)
            flag = 0
        elif token == 'высокий':
            ret.append(4)
        elif token == 'не':
            
            flag = 5
        else:
            flag == 0
    return ret

text = "не употребительный интеренсный спирт"
print("Введите терм: ")
text = input()

tokens = nltk.word_tokenize(text.lower(), language="russian")
tokens = [token for token in tokens if token not in punctuation]

# В вопросе спрашивают про спирт?
print('Разбитые токены:', tokens)

if not take_murav_from_tokens(tokens):
    perror("В вопросе нет слова 'спирт'")
    exit(1)

mass = find_terms(tokens)
terms = mass_to_feature(mass)

print('Найденные термы: ')
for i in range(len(terms) - 1):
    print('>   ' + terms[i] + '(%d), '%mass[i])
print(terms)
print(mass)
print('>   ' + terms[len(terms) - 1] + '(%d)'%mass[len(terms) - 1])

print('Подпадающие под условия участки:')
for uch, mur in conties_murav.items():
    if mur in mass:
        print(uch + ' с ' + mass_to_feature([mur])[0] + '(%d)'%mur)