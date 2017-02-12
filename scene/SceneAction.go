// scene provides data structures and adheres to the AvvaOscMessageBuilder interface
package scene

// SceneActions provide a string to be using in a valid avva.studio scene OSC message
type SceneAction interface {
	actionName() string
}
