// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	starwarsv1 "github.com/okontajneroch/starwars/api/starwars/v1"
	scheme "github.com/okontajneroch/starwars/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// StarfightersGetter has a method to return a StarfighterInterface.
// A group's client should implement this interface.
type StarfightersGetter interface {
	Starfighters(namespace string) StarfighterInterface
}

// StarfighterInterface has methods to work with Starfighter resources.
type StarfighterInterface interface {
	Create(ctx context.Context, starfighter *starwarsv1.Starfighter, opts metav1.CreateOptions) (*starwarsv1.Starfighter, error)
	Update(ctx context.Context, starfighter *starwarsv1.Starfighter, opts metav1.UpdateOptions) (*starwarsv1.Starfighter, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, starfighter *starwarsv1.Starfighter, opts metav1.UpdateOptions) (*starwarsv1.Starfighter, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*starwarsv1.Starfighter, error)
	List(ctx context.Context, opts metav1.ListOptions) (*starwarsv1.StarfighterList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *starwarsv1.Starfighter, err error)
	StarfighterExpansion
}

// starfighters implements StarfighterInterface
type starfighters struct {
	*gentype.ClientWithList[*starwarsv1.Starfighter, *starwarsv1.StarfighterList]
}

// newStarfighters returns a Starfighters
func newStarfighters(c *StarwarsV1Client, namespace string) *starfighters {
	return &starfighters{
		gentype.NewClientWithList[*starwarsv1.Starfighter, *starwarsv1.StarfighterList](
			"starfighters",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *starwarsv1.Starfighter { return &starwarsv1.Starfighter{} },
			func() *starwarsv1.StarfighterList { return &starwarsv1.StarfighterList{} },
		),
	}
}
