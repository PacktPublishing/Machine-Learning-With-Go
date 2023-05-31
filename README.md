


# Machine Learning With Go
This is the code repository for [Machine Learning With Go](https://www.packtpub.com/big-data-and-business-intelligence/machine-learning-go?utm_source=github&utm_medium=repository&utm_campaign=9781785882104), published by [Packt](https://www.packtpub.com/?utm_source=github). It contains all the supporting project files necessary to work through the book from start to finish.
## About the Book
The mission of this book is to turn readers into productive, innovative data analysts who leverage Go to build robust and valuable applications. To this end, the book clearly introduces the technical aspects of building predictive models in Go, but it also helps the reader understand how machine learning workflows are being applied in real-world scenarios.

## Instructions and Navigation
All of the code is organized into folders. Each folder starts with a number followed by the application name. For example, Chapter02.



The code will look like the following:
```
// Create a new matrix a.
a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
// Compute and output the transpose of the matrix.
ft := mat.Formatted(a.T(), mat.Prefix(" "))
fmt.Printf("a^T = %v\n\n", ft)
// Compute and output the determinant of a.
deta := mat.Det(a)
fmt.Printf("det(a) = %.2f\n\n", deta)
// Compute and output the inverse of a.
aInverse := mat.NewDense(0, 0, nil)
if err := aInverse.Inverse(a); err != nil {
log.Fatal(err)
}
fi := mat.Formatted(aInverse, mat.Prefix(" "))
fmt.Printf("a^-1 = %v\n\n", fi)
```

To run the examples in this book and experiment with the techniques covered in the book,
you will generally need the following:
Access to a bash-like shell.
A complete Go environment including Go, an editor, and related default or
custom environment variables defined. You can, for example, follow this guide
at https://www.goinggo.net/2016/05/installing-go-and-your-workspace.htm
l.
Various Go dependencies. These can be obtained as they are needed via `dep
ensure -v`
Then, to run the examples related to some of the advanced topics, such as data pipelining
and deep learning, you will need a few additional things:
An installation or deployment of Pachyderm. You can follow these docs to get
Pachyderm up and running locally or in the
cloud, http://pachyderm.readthedocs.io/en/latest/.
A working Docker installation
(https://www.docker.com/community-edition#/download).
An installation of TensorFlow. To install TensorFlow locally, you can follow this
guide at https://www.tensorflow.org/install/.

## Related Products
* [Mastering Machine Learning with scikit-learn - Second Edition](https://www.packtpub.com/big-data-and-business-intelligence/mastering-machine-learning-scikit-learn-second-edition?utm_source=github&utm_medium=repository&utm_campaign=9781788299879)

* [Mastering Machine Learning with scikit-learn](https://www.packtpub.com/big-data-and-business-intelligence/mastering-machine-learning-scikit-learn?utm_source=github&utm_medium=repository&utm_campaign=9781783988365)

* [Machine Learning with JavaScript](https://www.packtpub.com/big-data-and-business-intelligence/machine-learning-javascript?utm_source=github&utm_medium=repository&utm_campaign=9781787280199)

### Download a free PDF

 <i>If you have already purchased a print or Kindle version of this book, you can get a DRM-free PDF version at no cost.<br>Simply click on the link to claim your free PDF.</i>
<p align="center"> <a href="https://packt.link/free-ebook/9781785882104">https://packt.link/free-ebook/9781785882104 </a> </p>