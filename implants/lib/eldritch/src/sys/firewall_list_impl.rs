use starlark::const_frozen_string;
use starlark::values::Value;
use starlark::collections::SmallMap;
use starlark::values::Heap;
use starlark::values::dict::Dict;
use anyhow::{anyhow, Result};

#[cfg(target_os = "windows")]
use winfw::get_fw_rules;

#[cfg(target_os = "windows")]
fn list_firewall_rules_windows(starlark_heap: &Heap) -> Result<Vec<Dict>> {
    let rules = get_fw_rules().map_err(|e| anyhow!("Failed to get rules: {}", e))?;
    let mut rules_vec = Vec::new();

    for rule in rules {
        let rule_map = SmallMap::new();
        let mut dict_rule = Dict::new(rule_map);
        dict_rule.insert_hashed(const_frozen_string!("name").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.name).to_value());
        dict_rule.insert_hashed(const_frozen_string!("description").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.description).to_value());
        dict_rule.insert_hashed(const_frozen_string!("app_name").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.app_name).to_value());
        dict_rule.insert_hashed(const_frozen_string!("service_name").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.service_name).to_value());
        dict_rule.insert_hashed(const_frozen_string!("protocol").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.protocol.to_string()).to_value());
        dict_rule.insert_hashed(const_frozen_string!("icmp_type").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.icmp_type).to_value());
        dict_rule.insert_hashed(const_frozen_string!("local_ports").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.local_ports).to_value());
        dict_rule.insert_hashed(const_frozen_string!("remote_ports").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.remote_ports).to_value());
        dict_rule.insert_hashed(const_frozen_string!("local_addresses").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.local_adresses).to_value());
        dict_rule.insert_hashed(const_frozen_string!("remote_addresses").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.remote_addresses).to_value());
        dict_rule.insert_hashed(const_frozen_string!("profile1").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.profile1).to_value());
        dict_rule.insert_hashed(const_frozen_string!("profile2").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.profile2).to_value());
        dict_rule.insert_hashed(const_frozen_string!("profile3").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.profile3).to_value());
        dict_rule.insert_hashed(const_frozen_string!("direction").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.direction.to_string()).to_value());
        dict_rule.insert_hashed(const_frozen_string!("action").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.action.to_string()).to_value());
        dict_rule.insert_hashed(const_frozen_string!("interface_types").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.interface_types).to_value());
        dict_rule.insert_hashed(const_frozen_string!("interfaces").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.interfaces).to_value());
        dict_rule.insert_hashed(const_frozen_string!("enabled").to_value().get_hashed()?, Value::new_bool(rule.enabled));
        dict_rule.insert_hashed(const_frozen_string!("edge_traversal").to_value().get_hashed()?, Value::new_bool(rule.edge_traversal));
        dict_rule.insert_hashed(const_frozen_string!("grouping").to_value().get_hashed()?, starlark_heap.alloc_str(&rule.grouping).to_value());
        rules_vec.push(dict_rule);
    }
    Ok(rules_vec)
}

#[cfg(not(target_os = "windows"))]
fn list_firewall_rules_windows(_starlark_heap: &Heap) -> Result<Vec<Dict>> {
    Ok(Vec::new())
}


#[cfg(target_os = "linux")]
fn list_firewall_rules_linux(starlark_heap: &Heap) -> Result<Vec<Dict>> {
    Ok(Vec::new())
}

#[cfg(not(target_os = "linux"))]
fn list_firewall_rules_linux(_starlark_heap: &Heap) -> Result<Vec<Dict>> {
    Ok(Vec::new())
}

#[cfg(target_os = "macos")]
fn list_firewall_rules_macos(_starlark_heap: &Heap) -> Result<Vec<Dict>> {
    Ok(Vec::new())
}

#[cfg(not(target_os = "macos"))]
fn list_firewall_rules_macos(_starlark_heap: &Heap) -> Result<Vec<Dict>> {
    Ok(Vec::new())
}

fn firewall_list(starlark_heap: &Heap) -> Result<Dict> {
    let res_map = SmallMap::new();
    let mut res_dict = Dict::new(res_map);
    res_dict.insert_hashed(const_frozen_string!("windows").to_value().get_hashed()?, starlark_heap.alloc(list_firewall_rules_windows(starlark_heap)?));
    res_dict.insert_hashed(const_frozen_string!("linux").to_value().get_hashed()?, starlark_heap.alloc(list_firewall_rules_linux(starlark_heap)?));
    res_dict.insert_hashed(const_frozen_string!("macos").to_value().get_hashed()?, starlark_heap.alloc(list_firewall_rules_macos(starlark_heap)?));
    Ok(res_dict)
}

#[cfg(test)]
mod tests {
    
}