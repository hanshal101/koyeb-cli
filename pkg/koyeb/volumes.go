package koyeb

import (
	"github.com/spf13/cobra"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/errors"
)

func NewVolumeCmd() *cobra.Command {
	h := NewVolumeHandler()
	_ = h

	volumeCmd := &cobra.Command{
		Use:     "volumes ACTION",
		Aliases: []string{"vol", "volume"},
		Short:   "Manage persistent volumes",
	}

	createVolumeCmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Create a new volume",
		Args:  cobra.ExactArgs(1),
		RunE: WithCLIContext(func(ctx *CLIContext, cmd *cobra.Command, args []string) error {
			req := koyeb.NewCreatePersistentVolumeRequestWithDefaults()

			req.SetName(args[0])

			region, err := cmd.Flags().GetString("region")
			if err != nil {
				return err
			}
			if region == "" {
				return &errors.CLIError{
					What:     "Missing flag",
					Why:      "Missing --region flag",
					Solution: "Please specify the region of the volume",
				}
			}
			req.SetRegion(region)

			// TODO: use a flag for the volume type when/if we support more than one
			req.SetVolumeType(koyeb.PERSISTENTVOLUMEBACKINGSTORE_LOCAL_BLK)

			size, err := cmd.Flags().GetInt64("size")
			if err != nil {
				return err
			}
			req.SetMaxSize(size)

			readOnly, err := cmd.Flags().GetBool("read-only")
			if err != nil {
				return err
			}
			req.SetReadOnly(readOnly)

			return h.Create(ctx, cmd, args, req)
		}),
	}
	createVolumeCmd.Flags().String("region", "was", "Region of the volume")
	createVolumeCmd.Flags().Int64("size", 20, "Size of the volume in GB")
	createVolumeCmd.Flags().Bool("read-only", false, "Force the volume to be read-only")
	volumeCmd.AddCommand(createVolumeCmd)

	getVolumeCmd := &cobra.Command{
		Use:   "get NAME",
		Short: "Get a volume",
		Args:  cobra.ExactArgs(1),
		RunE: WithCLIContext(func(ctx *CLIContext, cmd *cobra.Command, args []string) error {
			return h.Get(ctx, cmd, args)
		}),
	}
	volumeCmd.AddCommand(getVolumeCmd)

	listVolumeCmd := &cobra.Command{
		Use:   "list",
		Short: "List volumes",
		RunE: WithCLIContext(func(ctx *CLIContext, cmd *cobra.Command, args []string) error {
			return h.List(ctx, cmd, args)
		}),
	}
	volumeCmd.AddCommand(listVolumeCmd)

	updateVolumeCmd := &cobra.Command{
		Use:   "update NAME",
		Short: "Update a volume",
		Args:  cobra.ExactArgs(1),
		RunE: WithCLIContext(func(ctx *CLIContext, cmd *cobra.Command, args []string) error {
			req := koyeb.NewUpdatePersistentVolumeRequestWithDefaults()

			name, _ := cmd.Flags().GetString("name")
			if name != "" {
				req.SetName(name)
			}

			size, err := cmd.Flags().GetInt64("size")
			if err != nil {
				return err
			}
			req.SetMaxSize(size)

			return h.Update(ctx, cmd, args, req)
		}),
	}
	updateVolumeCmd.Flags().String("name", "", "Change the volume name")
	updateVolumeCmd.Flags().String("size", "", "Increase the volume size")
	volumeCmd.AddCommand(updateVolumeCmd)

	deleteVolumeCmd := &cobra.Command{
		Use:   "delete NAME",
		Short: "Delete a volume",
		Args:  cobra.ExactArgs(1),
		RunE: WithCLIContext(func(ctx *CLIContext, cmd *cobra.Command, args []string) error {
			return h.Delete(ctx, cmd, args)
		}),
	}
	volumeCmd.AddCommand(deleteVolumeCmd)

	return volumeCmd
}

func NewVolumeHandler() *VolumeHandler {
	return &VolumeHandler{}
}

type VolumeHandler struct {
}

func ResolveVolumeArgs(ctx *CLIContext, val string) (string, error) {
	volumeMapper := ctx.Mapper.Volume()
	id, err := volumeMapper.ResolveID(val)
	if err != nil {
		return "", err
	}
	return id, nil
}
