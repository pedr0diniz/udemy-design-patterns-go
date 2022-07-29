package composite

import "fmt"

type NeuronInterface interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}

	return result
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{
		Neurons: make([]Neuron, count),
	}
}

// Both Neuron and NeuronLayer implement the NeuronInterface
// Therefore, they can interact with each other
func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func NeuralNetworks() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)

	fmt.Printf("neuron1: %+v\n", neuron1)
	fmt.Printf("neuron2: %+v\n", neuron2)
	fmt.Printf("layer1: %+v\n", layer1)
	fmt.Printf("layer2: %+v\n", layer2)
}
