package conn

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Override struct {
	ServiceID string
	Address   string
	Insecure  bool
}

// ParseOverrides parses value, usually env.
// String is a comma `,` separated list of `ID=ADDRESS` pairs with
// possible addition of `;insecure` to the address.
func ParseOverrides(value string) ([]Override, error) {
	pairs := strings.Split(value, ",")
	overrides := make([]Override, 0, len(pairs))
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		parts := strings.Split(pair, "=")
		const expected = 2
		if len(parts) < expected {
			return nil, fmt.Errorf("missing %q in %q", "=", pair)
		}
		if len(parts) > expected {
			return nil, fmt.Errorf("unexpected second %q in %q", "=", pair)
		}

		serviceID := strings.TrimSpace(parts[0])
		address, option, _ := strings.Cut(parts[1], ";")
		address, option = strings.TrimSpace(address), strings.TrimSpace(option)
		override := Override{
			ServiceID: serviceID,
			Address:   address,
			Insecure:  false,
		}
		switch option {
		case "":
		case "insecure":
			override.Insecure = true
		default:
			return nil, fmt.Errorf("unknown option %q", option)
		}

		overrides = append(overrides, override)
	}
	return overrides, nil
}

// ParseResolverAndDialOptions parses value using ParseOverrides.
// String is a comma `,` separated list of `ID=ADDRESS[;insecure]` pairs.
// If an ID ends with `*`, [PrefixResolver] is returned, and [SingleResolver] otherwise.
// For each insecure override respective dial option is returned.
func ParseResolverAndDialOptions(value string) (Resolver, []grpc.DialOption, error) {
	overrides, err := ParseOverrides(value)
	if err != nil {
		return nil, nil, err
	}

	var resolvers []Resolver
	for _, o := range overrides {
		resolvers = append(resolvers, NewBasicResolver(o.ServiceID, o.Address))
	}

	var resultResolver Resolver
	if len(resolvers) == 1 {
		resultResolver = resolvers[0]
	} else {
		resultResolver = NewResolversChain(resolvers...)
	}

	dialOptions := make(map[string]grpc.DialOption, len(overrides))
	for _, o := range overrides {
		if o.Insecure {
			opt := grpc.WithTransportCredentials(insecure.NewCredentials())
			dialOptions[o.Address] = WithAddressDialOptions(Address(o.Address), opt)
		}
	}

	return resultResolver, maps.Values(dialOptions), err
}

// ParseResolver parses value, usually env, and creates resolver.
// String is a comma `,` separated list of `ID=ADDRESS[;insecure]` pairs.
// If an ID ends with `*`, [PrefixResolver] is returned, and [SingleResolver] otherwise.
func ParseResolver(value string) (Resolver, error) {
	resolver, _, err := ParseResolverAndDialOptions(value)
	return resolver, err
}
