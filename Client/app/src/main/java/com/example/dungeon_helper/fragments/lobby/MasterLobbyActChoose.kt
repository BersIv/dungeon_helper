package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MasterLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentMasterLobbyActChooseBinding

class MasterLobbyActChoose : Fragment() {

    companion object {
        fun newInstance() = MasterLobbyActChoose()
    }

    private lateinit var viewModel: MasterLobbyActChooseViewModel
    private var _binding: FragmentMasterLobbyActChooseBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val masterLobbyActChooseViewModel = ViewModelProvider(this)[MasterLobbyActChooseViewModel::class.java]
        _binding = FragmentMasterLobbyActChooseBinding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textPlay
        masterLobbyActChooseViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val backBtn = binding.backBtn
        val hpBtn = binding.hpBtn
        val lvlBtn = binding.lvlBtn
        val expBtn = binding.expBtn
        hpBtn.setOnClickListener {
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyActChoose_to_masterLobbyHp)
        }
        expBtn.setOnClickListener {
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyActChoose_to_masterLobbyExp)
        }
        backBtn.setOnClickListener {
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyActChoose_to_masterLobbyChoose)
        }
    }

}