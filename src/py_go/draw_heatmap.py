# @author: code_king
# @date: 2024/1/11 21:00

import matplotlib.pyplot as plt
import numpy as np
from matplotlib.colors import ListedColormap

# 创建一个随机矩阵作为示例数据
# matrix_data = np.random.randint(0, 10, size=(6,6))
matrix_data = np.array([[0., 1., 0., 0., 0., 1.],
                        [1., 0., 1., 0., 0., 0.],
                        [0., 1., 0., 1., 0., 0.],
                        [0., 0., 1., 0., 1., 0.],
                        [0., 0., 0., 1., 0., 1.],
                        [1., 0., 0., 0., 1., 0.]])

full_step_matrix = np.array([[0., 1., 2., 3., 2., 1.],
                             [1., 0., 1., 2., 3., 2.],
                             [2., 1., 0., 1., 2., 3.],
                             [3., 2., 1., 0., 1., 2.],
                             [2., 3., 2., 1., 0., 1.],
                             [1., 2., 3., 2., 1., 0.]])

step_2_matrix=np.array([[0., 1., 2., 0., 2., 1.],
          [1., 0., 1., 2., 0., 2.],
          [2., 1., 0., 1., 2., 0.],
          [0., 2., 1., 0., 1., 2.],
          [2., 0., 2., 1., 0., 1.],
          [1., 2., 0., 2., 1., 0.]])

matrix_data = step_2_matrix

# 测试-----------
# matrix_data = np.zeros((6, 6), dtype=int)
# matrix_data[0][1] = 1
# matrix_data[1][0] = 1
# matrix_data[0][5] = 1
# matrix_data[5][0] = 1
# 测试-----------


# 绘制热力图，设置 origin='lower'，颜色 plasma，viridis，inferno，cubehelix,copper
# bicubic,nearest,bilinear,spline16,spline36,hanning,hamming,hermite,kaiser,quadric,catrom
# print(plt.colormaps())
# 创建自定义的 Blues 颜色映射
# custom_blues = ListedColormap(['#eff3ff', '#bdd7e7'])
# custom_blues = ListedColormap(['#eff3ff', '#bdd7e7', '#6baed6', '#3182bd', '#08519c'])
# def generate_colors(num_colors):
#     base_color='Pastel1_r'
#     # 选择一个颜色映射，这里使用 Blues
#     base_cmap = plt.colormaps.get_cmap(base_color)
#     # 生成渐变颜色
#     new_colors = base_cmap(np.linspace(0, 1, num_colors))
#     # 创建新的颜色映射
#     new_cmap = LinearSegmentedColormap.from_list('custom_cmap', new_colors, N=num_colors)
#     return new_cmap
#
# # 示例：生成颜色映射，根据列表长度确定颜色数量,到时候根据最大步长确定
# your_list = [1, 2]
# num_colors = len(your_list)
# custom_cmap = generate_colors(num_colors)

# 比较合适的颜色：Pastel1_r
# colorList=['#eff3ff', '#bdd7e7', '#6baed6', '#3182bd', '#08519c']
colorList = ['#eff3ff', '#bdd7e7','#6baed6']
custom_blues = ListedColormap(colorList[:2])
plt.imshow(matrix_data, cmap=custom_blues, interpolation='nearest', origin='lower')

# plt.imshow(matrix_data,cmap="Blues",  interpolation='nearest', origin='lower')

# 在每个矩阵元素上添加数字标签
for i in range(matrix_data.shape[0]):
    for j in range(matrix_data.shape[1]):
        plt.text(j, i, f'{matrix_data[i, j]:.0f}', ha='center', va='center', color='black')

# 获取当前的轴对象
ax = plt.gca()

# 显示刻度数字，同时隐藏刻度线
ax.set_xticks(np.arange(matrix_data.shape[1]))
ax.set_yticks(np.arange(matrix_data.shape[0]))

# 将坐标轴显示在上方
ax.xaxis.tick_top()
ax.xaxis.set_label_position('top')
ax.yaxis.tick_left()
ax.yaxis.set_label_position('left')
ax.invert_yaxis()

# 隐藏刻度线，同时显示坐标轴标签
ax.tick_params(axis='both', which='both', bottom=False, top=False, left=False, right=False,
               labelbottom=False, labeltop=True, labelleft=True, labelright=False)

# 显示热力图
plt.show()
