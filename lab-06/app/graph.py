import matplotlib.pyplot as plt
fig1 = plt.figure(figsize=(10, 7))
plot = fig1.add_subplot()
values = [1, 2, 3, 4, 5, 6, 7]
very_short = [1, 5/6, 1/6, 0, 0, 0, 0]
short = [2/6, 1, 3/6, 1/6, 0, 0, 0]
average = [0, 3/6, 1, 1/6, 0, 0, 0]
high = [0, 0, 0, 1/6, 1, 3/6, 2/6]
very_high = [0, 0, 0, 0, 2/6, 1, 4/6]

def expert_opinion(feature):
    if feature == 'легкий':
        return 1
    elif feature == '':
        return 2
    elif feature == '':
        return 3
    elif feature == '':
        return 4
    elif feature == '':
        return 5
    return None

plot.plot(values, very_short, label = "«легкий»")
plot.plot(values, short, ":", label="«средний»", marker="o")
plot.plot(values, average, "--", label="интеренсный", marker="v")
plot.plot(values, high, "--", label="«высокий»", marker="^")
plot.plot(values, very_high, "--", label="«не употребительный»", marker="s")

plt.legend()
plt.grid()
plt.title("Графики функций принадлежности числовых значений \nпеременной термам, описывающим группы значений лингвистической переменной")
plt.ylabel("μ")
plt.xlabel("количество спирта: 10^x")

plt.show()