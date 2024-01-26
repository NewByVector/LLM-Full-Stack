### 原理

$$
\begin{aligned}
&\text { Recap: } \hat{y}=\sigma\left(w^{T} x+b\right), \sigma(z)=\frac{1}{1+e^{-z}} \\
&J(w, b)=\frac{1}{m} \sum_{i=1}^{m} \mathcal{L}\left(\hat{y}^{(i)}, y^{(i)}\right)=-\frac{1}{m} \sum_{i=1}^{m} y^{(i)} \log \hat{y}^{(i)}+\left(1-y^{(i)}\right) \log \left(1-\hat{y}^{(i)}\right)
\end{aligned}
$$


![image.png](./image/07.png)
结合上面公式和图，其实成本函数是一个关于 w 和 b 的函数，神经网络学习的目的就是找到一个合适的 w 和 b，使得 J(w, b) 最小，可以通过梯度下降算法来实现。

![image.png](./image/08.png)

梯度下降算法的原理就是不断修改 w，不断逼近 J(w, b) 的最低点。可以通过下面公式来修改 w 值：

$$w^{\prime}=w-r^{\star} d w$$

其中 r 是学习率，dw 是 w 关于损失函数 J(w, b) 的偏导函数，就是上图的: $\frac{d J(\omega, b)}{d \omega}$

### 向量化

### 代码实现