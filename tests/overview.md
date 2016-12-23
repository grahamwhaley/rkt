# KVM functional tests overview

running QEMU tests on my F24, we got:
 - 81 tests passed
 - 3 tests FAIL:
    - TestGCAfterUnmount
    - TestNetDefaultsNetNS
    - TestNonRootReadInfo
 - 1 test timed out (panic)
    - TestSocketActivation (after 1h0m)

| File | kvm enbabled | lkvm ok | qemu ok | notes |
| ---- | ------------ | ------- | ------- | ----- |
| rkt_ace_validator_test.go | disabled | unk | unk  | awaiting test |
| rkt_api_service_bench_test.go | | | | |
| rkt_api_service_nspawn_test.go | !enabled | | | |
| rkt_api_service_test.go | | | | |
| rkt_app_isolator_nspawn_test.go | !enabled | | | |
| rkt_app_isolator_test.go | | | | |
| rkt_app_sandbox_test.go | disabled | | | |
| rkt_auth_test.go | | | | |
| rkt_caps_kvm_test.go | | | | |
| rkt_caps_nspawn_test.go | !enabled | | | |
| rkt_caps_test.go | | | | |
| rkt_cat_manifest_test.go | | | | |
| rkt_config_test.go | | | | |
| rkt_devices_test.go | !enabled | | | |
| rkt_diagnostic_test.go | | | | |
| rkt_dns_test.go | | | | |
| rkt_env_test.go | | | | |
| rkt_error_output_test.go | | | | |
| rkt_etc_hosts_test.go | | | | |
| rkt_exec_test.go | | | | |
| rkt_exit_test.go | !enabled | | | |
| rkt_export_container_test.go | | | | |
| rkt_export_fly_test.go | !enabled | | | |
| rkt_export_test.go | | | | |
| rkt_fetch_test.go | | | | |
| rkt_flavor_coreos.go | !enabled | | | |
| rkt_flavor_default.go | disabled | | | |
| rkt_flavor_fly.go | !enabled | | | |
| rkt_flavor.go | | | | |
| rkt_flavor_host.go | !enabled | | | |
| rkt_flavor_kvm.go | | | | |
| rkt_flavor_src.go | !enabled | | | |
| rkt_fly_test.go | !enabled | | | |
| rkt_gc_nspawn_test.go | !enabled | | | |
| rkt_gc_test.go | | | | |
| rkt_hostname_test.go | | | | |
| rkt_image_cat_manifest_test.go | | | | |
| rkt_image_dependencies_test.go | | | | |
| rkt_image_export_test.go | | | | |
| rkt_image_extract_test.go | | | | |
| rkt_image_gc_test.go | | | | |
| rkt_image_list_test.go | | | | |
| rkt_image_render_test.go | | | | |
| rkt_image_rm_test.go | | | | |
| rkt_interactive_test.go | | | | |
| rkt_journal_test.go | | | | |
| rkt_list_test.go | | | | |
| rkt_metadata_service_test.go | | | | |
| rkt_mountns_test.go | | | | |
| rkt_mount_test.go | !enabled | | | |
| rkt_mtab_test.go | | | | |
| rkt_net_kvm_test.go | | | | |
| rkt_net_nspawn_test.go | !enabled | | | |
| rkt_net_test.go | | | | |
| rkt_no_new_privs_test.go | !enabled | | | |
| rkt_non_root_test.go | | | | |
| rkt_oom_score_adj_test.go | !enabled | | | |
| rkt_os_arch_test.go | | | | |
| rkt_paths_test.go | !enabled | | | |
| rkt_pid_file_kvm_test.go | | | | |
| rkt_pid_file_nspawn_test.go | !enabled | | | |
| rkt_pid_file_test.go | | | | |
| rkt_prepare_test.go | | | | |
| rkt_rm_nspawn_test.go | | | | |
| rkt_rm_test.go | | | | |
| rkt_root_commands_test.go | | | | |
| rkt_run_pod_manifest_test.go | !enabled | | | |
| rkt_run_test.go | | | | |
| rkt_run_user_group_test.go | !enabled | | | |
| rkt_seccomp_test.go | !enabled | | | |
| rkt_service_file_test.go | | | | |
| rkt_socket_activation_test.go | | | | |
| rkt_socket_proxyd_test.go | | | | |
| rkt_stage1_loading_test.go | | | | |
| rkt_stop_test.go | | | | |
| rkt_supplementary_gids_test.go | | | | |
| rkt_tests.go | | | | |
| rkt_trust_test.go | | | | |
| rkt_userns_test.go | !enabled | | | |
| rkt_volume_mount_fly_test.go | !enabled | | | |
| rkt_volume_mount_generic_test.go | | | | |
| rkt_volume_mount_test.go | | | | |
| rkt_volume_nspawn_test.go | !enabled | | | |
| rkt_volume_test.go | | | | |
