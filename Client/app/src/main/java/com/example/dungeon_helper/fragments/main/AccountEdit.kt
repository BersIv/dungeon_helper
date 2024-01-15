package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentAccountEditBinding

class AccountEdit : Fragment() {

    companion object {
        fun newInstance() = AccountEdit()
    }

    private lateinit var viewModel: AccountEditViewModel

    private var _binding: FragmentAccountEditBinding? = null
    private val binding get() = _binding!!
    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val activity = requireActivity() as MainActivity
        val navView = activity.getNavView()
        val menu = navView.menu
        val menuItem1 = menu.findItem(R.id.navigation_info)
        val menuItem2 = menu.findItem(R.id.navigation_character)
        val menuItem3 = menu.findItem(R.id.navigation_lobby)
        val menuItem4 = menu.findItem(R.id.navigation_account)
        menuItem1.isEnabled = false
        menuItem2.isEnabled = false
        menuItem3.isEnabled = false
        menuItem4.isEnabled = false
       val accountEditViewModel = ViewModelProvider(this)[AccountEditViewModel::class.java]
        _binding = FragmentAccountEditBinding.inflate(inflater, container, false)
        val root: View =  binding.root
        val textView: TextView = binding.textAccountEdit
        accountEditViewModel.text.observe(viewLifecycleOwner)
        {
            textView.text = it
        }
        return root

    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
        val activity = requireActivity() as MainActivity
        val navView = activity.getNavView()
        val menu = navView.menu
        val menuItem1 = menu.findItem(R.id.navigation_info)
        val menuItem2 = menu.findItem(R.id.navigation_character)
        val menuItem3 = menu.findItem(R.id.navigation_lobby)
        val menuItem4 = menu.findItem(R.id.navigation_account)
        menuItem1.isEnabled = true
        menuItem2.isEnabled = true
        menuItem3.isEnabled = true
        menuItem4.isEnabled = true
    }

    override fun onStart() {
        super.onStart()
        val exAccBtn = binding.exAccBtn
        val exEditBtn = binding.exEditBtn
        val changePwdBtn = binding.changePwdBtn

        exAccBtn.setEnabled(false)
        changePwdBtn.setEnabled(false)
        exEditBtn.setOnClickListener{
            (activity as MainActivity).navController.navigate(R.id.action_accountEdit_to_navigation_account)
        }

    }

}


